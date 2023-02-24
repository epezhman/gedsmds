package mdsprocessor

import (
	"github.com/IBM/gedsmds/internal/keyvaluestore"
	"github.com/IBM/gedsmds/internal/logger"
	"github.com/IBM/gedsmds/protos/protos"
	"sync"
)

func InitProcessor(kvStore *keyvaluestore.KeyValueStore) *Processor {
	processor := &Processor{
		kvStore:               kvStore,
		objectSubscribersLock: &sync.RWMutex{},
		objectSubscribers:     map[string]*objectSubscriber{},
	}
	return processor
}

func (p *Processor) ObjectEventSubscription(subscription *protos.ObjectEventSubscription,
	stream protos.MetadataService_SubscribeObjectsServer) error {
	logger.InfoLogger.Println("got subscription subscriberID - objectId", subscription.SubscriberID, subscription.ObjectID)
	finished := make(chan bool)
	p.objectSubscribersLock.Lock()
	p.objectSubscribers[subscription.SubscriberID] = &objectSubscriber{
		stream:   stream,
		finished: finished,
	}
	p.objectSubscribersLock.Unlock()
	cntx := stream.Context()

	go func(sub *protos.ObjectEventSubscription) {
		for i := 0; i < 10; i++ {
			p.sendObjectPublication(sub.SubscriberID, &protos.ObjectEventSubscription{
				ObjectID: sub.SubscriberID,
			})
		}
	}(subscription)

	for {
		select {
		case <-finished:
			return nil
		case <-cntx.Done():
			return nil
		}
	}
}

func (p *Processor) sendObjectPublication(objectID string, publicationEvent *protos.ObjectEventSubscription) {
	p.objectSubscribersLock.RLock()
	streamer, ok := p.objectSubscribers[objectID]
	p.objectSubscribersLock.RUnlock()
	if !ok {
		return
	}
	publication := &protos.ObjectSubscription{
		ObjectID: publicationEvent.ObjectID,
	}
	if err := streamer.stream.Send(publication); err != nil {
		streamer.finished <- true
		logger.ErrorLogger.Println("Could not send the proposal response to the client " + objectID)
		p.objectSubscribersLock.Lock()
		delete(p.objectSubscribers, objectID)
		p.objectSubscribersLock.Unlock()
	}
	logger.InfoLogger.Println("sending publication subscriberID - objectID", objectID, publication.ObjectID)
}
