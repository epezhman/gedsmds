package keyvaluestore

import (
	"errors"
	"github.com/IBM/gedsmds/internal/config"
	"github.com/IBM/gedsmds/internal/keyvaluestore/db"
	"github.com/IBM/gedsmds/internal/logger"
	"github.com/IBM/gedsmds/protos/protos"
	"golang.org/x/exp/maps"
	"strings"
	"sync"
)

func InitKeyValueStoreService() *Service {
	kvStore := &Service{
		dbConnection: db.NewOperations(),

		ObjectStoreConfigsLock: &sync.RWMutex{},
		ObjectStoreConfigs:     map[string]*protos.ObjectStoreConfig{},

		BucketsLock: &sync.RWMutex{},
		Buckets:     map[string]*protos.Bucket{},

		// possible bottleneck if every item written in the same map
		ObjectsLock: &sync.RWMutex{},
		Objects:     map[string]map[string]*Object{},
	}
	go kvStore.populateCache()
	return kvStore
}

func (kv *Service) populateCache() {
	if config.Config.PersistentStorageEnabled && config.Config.RepopulateCacheEnabled {
		go kv.populateObjectStoreConfig()
		go kv.populateBuckets()
		go kv.populateObjects()
	}
}

func (kv *Service) populateObjectStoreConfig() {
	if allObjectStoreConfig, err := kv.dbConnection.GetAllObjectStoreConfig(); err != nil {
		logger.ErrorLogger.Println(err)
	} else {
		kv.ObjectStoreConfigsLock.Lock()
		for _, objectStoreConfig := range allObjectStoreConfig {
			kv.ObjectStoreConfigs[objectStoreConfig.Bucket] = objectStoreConfig
		}
		kv.ObjectStoreConfigsLock.Unlock()
	}
}

func (kv *Service) populateBuckets() {
	if allBuckets, err := kv.dbConnection.GetAllBuckets(); err != nil {
		logger.ErrorLogger.Println(err)
	} else {
		kv.BucketsLock.Lock()
		for _, bucket := range allBuckets {
			logger.InfoLogger.Println(bucket)
			kv.Buckets[bucket.Bucket] = bucket
		}
		kv.BucketsLock.Unlock()
	}
}

func (kv *Service) populateObjects() {
	if allObjects, err := kv.dbConnection.GetAllObjects(); err != nil {
		logger.ErrorLogger.Println(err)
	} else {
		kv.ObjectsLock.Lock()
		for _, object := range allObjects {
			if _, ok := kv.Objects[object.Id.Bucket]; !ok {
				kv.Objects[object.Id.Bucket] = map[string]*Object{}
			}
			kv.Objects[object.Id.Bucket][object.Id.Key] = &Object{object: object, path: kv.getNestedPath(object.Id)}
		}
		kv.ObjectsLock.Unlock()
	}
}

func (kv *Service) RegisterObjectStore(objectStore *protos.ObjectStoreConfig) error {
	kv.ObjectStoreConfigsLock.Lock()
	defer kv.ObjectStoreConfigsLock.Unlock()
	if _, ok := kv.ObjectStoreConfigs[objectStore.Bucket]; ok {
		return errors.New("config already exists")
	}
	kv.ObjectStoreConfigs[objectStore.Bucket] = objectStore
	if config.Config.PersistentStorageEnabled {
		kv.dbConnection.ObjectStoreConfigChan <- &db.OperationParams{
			ObjectStoreConfig: objectStore,
			Type:              db.PUT,
		}
	}
	return nil
}

func (kv *Service) ListObjectStores() (*protos.AvailableObjectStoreConfigs, error) {
	kv.ObjectStoreConfigsLock.RLock()
	defer kv.ObjectStoreConfigsLock.RUnlock()
	mappings := &protos.AvailableObjectStoreConfigs{Mappings: []*protos.ObjectStoreConfig{}}
	for _, objectStoreConfig := range kv.ObjectStoreConfigs {
		mappings.Mappings = append(mappings.Mappings, objectStoreConfig)
	}
	return mappings, nil
}

func (kv *Service) CreateBucket(bucket *protos.Bucket) error {
	kv.BucketsLock.Lock()
	defer kv.BucketsLock.Unlock()
	if _, ok := kv.Buckets[bucket.Bucket]; ok {
		return errors.New("bucket already exists")
	}
	kv.Buckets[bucket.Bucket] = bucket
	if config.Config.PersistentStorageEnabled {
		kv.dbConnection.BucketChan <- &db.OperationParams{
			Bucket: bucket,
			Type:   db.PUT,
		}
	}
	return nil
}

func (kv *Service) DeleteBucket(bucket *protos.Bucket) error {
	kv.BucketsLock.Lock()
	defer kv.BucketsLock.Unlock()
	if _, ok := kv.Buckets[bucket.Bucket]; !ok {
		return errors.New("bucket already deleted")
	}
	delete(kv.Buckets, bucket.Bucket)
	if config.Config.PersistentStorageEnabled {
		kv.dbConnection.BucketChan <- &db.OperationParams{
			Bucket: bucket,
			Type:   db.DELETE,
		}
	}
	return nil
}

func (kv *Service) ListBuckets() (*protos.BucketListResponse, error) {
	kv.BucketsLock.RLock()
	defer kv.BucketsLock.RUnlock()
	buckets := &protos.BucketListResponse{Results: []string{}}
	buckets.Results = append(buckets.Results, maps.Keys(kv.Buckets)...)
	return buckets, nil
}

func (kv *Service) LookupBucket(bucket *protos.Bucket) error {
	kv.BucketsLock.RLock()
	defer kv.BucketsLock.RUnlock()
	if _, ok := kv.Buckets[bucket.Bucket]; !ok {
		return errors.New("bucket does not exist")
	}
	return nil
}

func (kv *Service) LookupBucketByName(bucketName string) error {
	kv.BucketsLock.RLock()
	defer kv.BucketsLock.RUnlock()
	if _, ok := kv.Buckets[bucketName]; !ok {
		return errors.New("bucket does not exist")
	}
	return nil
}

func (kv *Service) CreateObject(object *protos.Object) error {
	if err := kv.LookupBucketByName(object.Id.Bucket); err != nil {
		logger.ErrorLogger.Println("bucket does not exist: ", object.Id.Bucket)
	}
	kv.ObjectsLock.Lock()
	defer kv.ObjectsLock.Unlock()
	if _, ok := kv.Objects[object.Id.Bucket][object.Id.Key]; ok {
		logger.InfoLogger.Println("object already exists %+v", object)
	}
	if _, ok := kv.Objects[object.Id.Bucket]; !ok {
		kv.Objects[object.Id.Bucket] = map[string]*Object{}
	}
	kv.Objects[object.Id.Bucket][object.Id.Key] = &Object{object: object, path: kv.getNestedPath(object.Id)}
	if config.Config.PersistentStorageEnabled {
		kv.dbConnection.ObjectChan <- &db.OperationParams{
			Object: object,
			Type:   db.PUT,
		}
	}
	return nil
}

func (kv *Service) UpdateObject(object *protos.Object) error {
	kv.ObjectsLock.Lock()
	defer kv.ObjectsLock.Unlock()
	if _, ok := kv.Objects[object.Id.Bucket]; !ok {
		kv.Objects[object.Id.Bucket] = map[string]*Object{}
	}
	kv.Objects[object.Id.Bucket][object.Id.Key] = &Object{object: object, path: kv.getNestedPath(object.Id)}
	if config.Config.PersistentStorageEnabled {
		kv.dbConnection.ObjectChan <- &db.OperationParams{
			Object: object,
			Type:   db.PUT,
		}
	}
	return nil
}

func (kv *Service) CreateOrUpdateObject(object *protos.Object) error {
	if err := kv.LookupBucketByName(object.Id.Bucket); err != nil {
		logger.ErrorLogger.Println("bucket does not exist: ", object.Id.Bucket)
	}
	kv.ObjectsLock.Lock()
	defer kv.ObjectsLock.Unlock()
	if _, ok := kv.Objects[object.Id.Bucket]; !ok {
		kv.Objects[object.Id.Bucket] = map[string]*Object{}
	}
	kv.Objects[object.Id.Bucket][object.Id.Key] = &Object{object: object, path: kv.getNestedPath(object.Id)}
	if config.Config.PersistentStorageEnabled {
		kv.dbConnection.ObjectChan <- &db.OperationParams{
			Object: object,
			Type:   db.PUT,
		}
	}
	return nil
}

func (kv *Service) DeleteObject(objectID *protos.ObjectID) error {
	kv.ObjectsLock.Lock()
	defer kv.ObjectsLock.Unlock()
	if _, ok := kv.Objects[objectID.Bucket][objectID.Key]; !ok {
		return errors.New("object already deleted")
	}
	delete(kv.Objects[objectID.Bucket], objectID.Key)
	if config.Config.PersistentStorageEnabled {
		kv.dbConnection.ObjectChan <- &db.OperationParams{
			Object: &protos.Object{
				Id: objectID,
			},
			Type: db.DELETE,
		}
	}
	return nil
}

func (kv *Service) DeleteObjectPrefix(objectID *protos.ObjectID) ([]*protos.Object, error) {
	var objects []*protos.Object
	// this will be slow, needs to be optimized
	kv.ObjectsLock.Lock()
	for key, object := range kv.Objects[objectID.Bucket] {
		if strings.HasPrefix(key, objectID.Key) {
			objects = append(objects, object.object)
			delete(kv.Objects, objectID.Key)
		}
	}
	kv.ObjectsLock.Unlock()
	if config.Config.PersistentStorageEnabled {
		for _, object := range objects {
			kv.dbConnection.ObjectChan <- &db.OperationParams{
				Object: object,
				Type:   db.DELETE,
			}
		}
	}
	return objects, nil
}

func (kv *Service) LookupObject(objectID *protos.ObjectID) (*protos.ObjectResponse, error) {
	kv.ObjectsLock.RLock()
	defer kv.ObjectsLock.RUnlock()
	if _, ok := kv.Objects[objectID.Bucket][objectID.Key]; !ok {
		return nil, errors.New("object does not exist")
	}
	return &protos.ObjectResponse{
		Result: kv.Objects[objectID.Bucket][objectID.Key].object,
	}, nil
}

func (kv *Service) ListObjects(objectListRequest *protos.ObjectListRequest) (*protos.ObjectListResponse, error) {
	objects := &protos.ObjectListResponse{Results: []*protos.Object{}, CommonPrefixes: []string{}}
	if objectListRequest.Prefix == nil || len(objectListRequest.Prefix.Bucket) == 0 {
		logger.InfoLogger.Println("bucket not set")
		return objects, nil
	}
	var delimiter string
	if objectListRequest.Delimiter != nil && *objectListRequest.Delimiter != 0 {
		delimiter = string(*objectListRequest.Delimiter)
	}
	tempCommonPrefixes := map[string]bool{}
	// needs to be optimized
	if len(delimiter) == 0 {
		kv.ObjectsLock.RLock()
		for key, object := range kv.Objects[objectListRequest.Prefix.Bucket] {
			if strings.HasPrefix(key, objectListRequest.Prefix.Key) {
				objects.Results = append(objects.Results, object.object)
			}
		}
		kv.ObjectsLock.RUnlock()
	} else {
		if len(objectListRequest.Prefix.Key) == 0 {
			kv.ObjectsLock.RLock()
			for _, object := range kv.Objects[objectListRequest.Prefix.Bucket] {
				if len(object.path) == 1 {
					objects.Results = append(objects.Results, object.object)
				} else if len(object.path) > 1 {
					tempCommonPrefixes[object.path[0]] = true
				}
			}
			kv.ObjectsLock.RUnlock()
		} else {
			prefixLength := len(strings.Split(objectListRequest.Prefix.Key, delimiter)) + 1
			kv.ObjectsLock.RLock()
			for key, object := range kv.Objects[objectListRequest.Prefix.Bucket] {
				if strings.HasPrefix(key, objectListRequest.Prefix.Key) {
					objects.Results = append(objects.Results, object.object)
					if len(object.path) == prefixLength {
						tempCommonPrefixes[key+object.path[prefixLength-1]] = true
					}
				}
			}
			kv.ObjectsLock.RUnlock()
		}
	}
	if len(tempCommonPrefixes) > 0 {
		for commonPrefix := range tempCommonPrefixes {
			objects.CommonPrefixes = append(objects.CommonPrefixes, commonPrefix+commonDelimiter)
		}
	}

	return objects, nil
}
