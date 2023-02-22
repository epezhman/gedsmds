package mdsprocessor

import (
	"github.com/IBM/gedsmds/internal/keyvaluestore"
	"github.com/IBM/gedsmds/protos/protos"
	"sync"
)

type objectSubscriber struct {
	stream   protos.MetadataService_SubscribeObjectsServer
	finished chan<- bool
}

type Processor struct {
	kvStore               *keyvaluestore.KeyValueStore
	objectSubscribersLock *sync.RWMutex
	objectSubscribers     map[string]*objectSubscriber
}
