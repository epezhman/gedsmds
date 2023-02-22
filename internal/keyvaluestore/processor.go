package keyvaluestore

import (
	"errors"
	"github.com/IBM/gedsmds/internal/keyvaluestore/db"
	"github.com/IBM/gedsmds/internal/logger"
	"github.com/IBM/gedsmds/protos/protos"
	"golang.org/x/exp/maps"
	"sync"
)

func InitKeyValueStore() *KeyValueStore {
	kvStore := &KeyValueStore{
		kvObjectStoreConfigLock:  &sync.RWMutex{},
		kvObjectStoreConfigMap:   map[string]*protos.ObjectStoreConfig{},
		kvObjectStoreConfigSlice: []*protos.ObjectStoreConfig{},
		dbConnection:             db.NewOperations(),
		kvBucketLock:             &sync.RWMutex{},
		kvBucket:                 map[string]*protos.Bucket{},
		kvObjectsLock:            &sync.RWMutex{},
		kvObjectsMap:             map[string]*protos.Object{},
	}
	return kvStore
}

func (kv *KeyValueStore) RegisterObjectStore(objectStore *protos.ObjectStoreConfig) error {
	kv.kvObjectStoreConfigLock.Lock()
	defer kv.kvObjectStoreConfigLock.Unlock()
	if _, ok := kv.kvObjectStoreConfigMap[objectStore.Bucket]; ok {
		return errors.New("config already exists")
	}
	kv.kvObjectStoreConfigMap[objectStore.Bucket] = objectStore
	kv.kvObjectStoreConfigSlice = append(kv.kvObjectStoreConfigSlice, objectStore)
	return nil
}

func (kv *KeyValueStore) ListObjectStores() (*protos.AvailableObjectStoreConfigs, error) {
	kv.kvObjectStoreConfigLock.RLock()
	defer kv.kvObjectStoreConfigLock.RUnlock()
	mappings := &protos.AvailableObjectStoreConfigs{Mappings: []*protos.ObjectStoreConfig{}}
	mappings.Mappings = append(mappings.Mappings, kv.kvObjectStoreConfigSlice...)
	return mappings, nil
}

func (kv *KeyValueStore) CreateBucket(bucket *protos.Bucket) error {
	kv.kvBucketLock.Lock()
	defer kv.kvBucketLock.Unlock()
	if _, ok := kv.kvBucket[bucket.Bucket]; ok {
		return errors.New("bucket already exists")
	}
	kv.kvBucket[bucket.Bucket] = bucket
	return nil
}

func (kv *KeyValueStore) DeleteBucket(bucket *protos.Bucket) error {
	kv.kvBucketLock.Lock()
	defer kv.kvBucketLock.Unlock()
	if _, ok := kv.kvBucket[bucket.Bucket]; !ok {
		return errors.New("bucket already deleted")
	}
	delete(kv.kvBucket, bucket.Bucket)
	return nil
}

func (kv *KeyValueStore) ListBuckets() (*protos.BucketListResponse, error) {
	kv.kvBucketLock.RLock()
	defer kv.kvBucketLock.RUnlock()
	buckets := &protos.BucketListResponse{Results: []string{}}
	buckets.Results = append(buckets.Results, maps.Keys(kv.kvBucket)...)
	return buckets, nil
}

func (kv *KeyValueStore) LookupBucket(bucket *protos.Bucket) error {
	kv.kvBucketLock.RLock()
	defer kv.kvBucketLock.RUnlock()
	if _, ok := kv.kvBucket[bucket.Bucket]; !ok {
		return errors.New("bucket does not exist")
	}
	return nil
}

func (kv *KeyValueStore) LookupBucketByName(bucketName string) error {
	kv.kvBucketLock.RLock()
	defer kv.kvBucketLock.RUnlock()
	if _, ok := kv.kvBucket[bucketName]; !ok {
		return errors.New("bucket does not exist")
	}
	return nil
}

func (kv *KeyValueStore) CreateObject(object *protos.Object) error {
	if err := kv.LookupBucketByName(object.Id.Bucket); err != nil {
		return err
	}
	kv.kvObjectsLock.Lock()
	defer kv.kvObjectsLock.Unlock()
	if _, ok := kv.kvObjectsMap[object.Id.Key]; ok {
		logger.InfoLogger.Println("object already exists")
	}
	kv.kvObjectsMap[object.Id.Key] = object
	return nil
}

func (kv *KeyValueStore) UpdateObject(object *protos.Object) error {
	kv.kvObjectsLock.Lock()
	defer kv.kvObjectsLock.Unlock()
	if _, ok := kv.kvObjectsMap[object.Id.Key]; ok {
		logger.InfoLogger.Println("object already exists")
	}
	kv.kvObjectsMap[object.Id.Key] = object
	return nil
}

func (kv *KeyValueStore) DeleteObject(objectID *protos.ObjectID) error {
	kv.kvObjectsLock.Lock()
	defer kv.kvObjectsLock.Unlock()
	if _, ok := kv.kvObjectsMap[objectID.Key]; !ok {
		return errors.New("object already deleted")
	}
	delete(kv.kvObjectsMap, objectID.Key)
	return nil
}

func (kv *KeyValueStore) DeleteObjectPrefix(_ *protos.ObjectID) error {
	return nil
}

func (kv *KeyValueStore) LookupObject(objectID *protos.ObjectID) (*protos.ObjectResponse, error) {
	kv.kvObjectsLock.RLock()
	defer kv.kvObjectsLock.RUnlock()
	if _, ok := kv.kvObjectsMap[objectID.Key]; !ok {
		return nil, errors.New("object does not exist")
	}
	return &protos.ObjectResponse{
		Result: kv.kvObjectsMap[objectID.Key],
	}, nil
}

func (kv *KeyValueStore) ListObjects() error {
	return nil
}
