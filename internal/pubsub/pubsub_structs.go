package pubsub

import (
	"github.com/IBM/gedsmds/internal/keyvaluestore"
	"github.com/IBM/gedsmds/protos/protos"
	"sync"
)

type SubscriberStream struct {
	stream   protos.MetadataService_SubscribeServer
	finished chan<- bool
}

type Service struct {
	kvStore *keyvaluestore.KeyValueStoreService

	bucketSubscribersLock   *sync.RWMutex
	bucketSubscriberStreams map[string]*SubscriberStream
	bucketSubscribers       map[string][]string

	objectSubscribersLock   *sync.RWMutex
	objectSubscriberStreams map[string]*SubscriberStream
	objectSubscribers       map[string][]string

	prefixSubscribersLock   *sync.RWMutex
	prefixSubscriberStreams map[string]*SubscriberStream
	prefixSubscribers       map[string][]string
}