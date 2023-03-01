package keyvaluestore

import (
	"errors"
	"github.com/IBM/gedsmds/internal/config"
	"github.com/IBM/gedsmds/internal/keyvaluestore/db"
	"github.com/IBM/gedsmds/internal/logger"
	"github.com/IBM/gedsmds/protos/protos"
	"golang.org/x/exp/maps"
	"sync"
)

func InitKeyValueStoreService() *Service {
	kvStore := &Service{
		dbConnection: db.NewOperations(),

		ObjectStoreConfigsLock: &sync.RWMutex{},
		ObjectStoreConfigs:     map[string]*protos.ObjectStoreConfig{},

		BucketsLock: &sync.RWMutex{},
		Buckets:     map[string]*protos.Bucket{},

		ObjectsLock: &sync.RWMutex{},
		Objects:     map[string]*protos.Object{},
	}
	go kvStore.populateCache()
	return kvStore
}

func (kv *Service) populateCache() {
	if config.Config.PersistentStorageEnabled {
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
			kv.Objects[kv.createObjectKey(object)] = object
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
		return err
	}
	kv.ObjectsLock.Lock()
	defer kv.ObjectsLock.Unlock()
	objectId := kv.createObjectKey(object)
	if _, ok := kv.Objects[objectId]; ok {
		logger.InfoLogger.Println("object already exists %+v", object)
	}
	kv.Objects[objectId] = object
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
	kv.Objects[kv.createObjectKey(object)] = object
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
	fakeObject := &protos.Object{
		Id: objectID,
	}
	objectId := kv.createObjectKey(fakeObject)
	if _, ok := kv.Objects[objectId]; !ok {
		return errors.New("object already deleted")
	}
	delete(kv.Objects, objectId)
	if config.Config.PersistentStorageEnabled {
		kv.dbConnection.ObjectChan <- &db.OperationParams{
			Object: fakeObject,
			Type:   db.DELETE,
		}
	}
	return nil
}

func (kv *Service) DeleteObjectPrefix(_ *protos.ObjectID) error {
	return nil
}

func (kv *Service) LookupObject(objectID *protos.ObjectID) (*protos.ObjectResponse, error) {
	kv.ObjectsLock.RLock()
	defer kv.ObjectsLock.RUnlock()
	objectId := kv.createObjectKey(&protos.Object{
		Id: objectID,
	})
	if _, ok := kv.Objects[objectId]; !ok {
		return nil, errors.New("object does not exist")
	}
	return &protos.ObjectResponse{
		Result: kv.Objects[objectId],
	}, nil
}

func (kv *Service) ListObjects() error {
	return nil
}
