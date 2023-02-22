package keyvaluestore

import (
	"github.com/IBM/gedsmds/internal/keyvaluestore/db"
	"github.com/IBM/gedsmds/protos/protos"
	"sync"
)

type KeyValueStore struct {
	kvObjectStoreConfigLock  *sync.RWMutex
	kvObjectStoreConfigMap   map[string]*protos.ObjectStoreConfig
	kvObjectStoreConfigSlice []*protos.ObjectStoreConfig
	dbConnection             *db.Operations
	kvBucketLock             *sync.RWMutex
	kvBucket                 map[string]*protos.Bucket
	kvObjectsLock            *sync.RWMutex
	kvObjectsMap             map[string]*protos.Object
}
