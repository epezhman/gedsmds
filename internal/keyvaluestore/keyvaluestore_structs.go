package keyvaluestore

import (
	"github.com/IBM/gedsmds/internal/keyvaluestore/db"
	"github.com/IBM/gedsmds/protos/protos"
	"sync"
)

const commonDelimiter = "/"

type Service struct {
	dbConnection *db.Operations

	ObjectStoreConfigsLock *sync.RWMutex
	ObjectStoreConfigs     map[string]*protos.ObjectStoreConfig

	BucketsLock *sync.RWMutex
	Buckets     map[string]*protos.Bucket

	ObjectsLock *sync.RWMutex
	Objects     map[string]map[string]*Object
}

type Object struct {
	path   []string
	object *protos.Object
}
