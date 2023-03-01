package pubsub

import (
	"errors"
	"github.com/IBM/gedsmds/internal/keyvaluestore"
	"github.com/IBM/gedsmds/internal/logger"
	"github.com/IBM/gedsmds/protos/protos"
	"sync"
)

func InitService(kvStore *keyvaluestore.Service) *Service {
	service := &Service{
		kvStore: kvStore,

		bucketSubscribersLock:   &sync.RWMutex{},
		bucketSubscriberStreams: map[string]*SubscriberStream{},
		bucketSubscribers:       map[string][]string{},

		objectSubscribersLock:   &sync.RWMutex{},
		objectSubscriberStreams: map[string]*SubscriberStream{},
		objectSubscribers:       map[string][]string{},

		prefixSubscribersLock:   &sync.RWMutex{},
		prefixSubscriberStreams: map[string]*SubscriberStream{},
		prefixSubscribers:       map[string][]string{},

		UpdatedObject: make(chan *protos.Object, channelBufferSize),
	}
	go service.runPubSubEventListeners()
	return service
}

func (s *Service) runPubSubEventListeners() {
	for {
		select {
		case object := <-s.UpdatedObject:
			bucket := &protos.Bucket{Bucket: object.Id.Bucket}
			go s.matchSubscriptions(&protos.SubscriptionEvent{
				SubscriptionType: protos.SubscriptionType_OBJECT,
				Key:              object.Id.Key,
				BucketID:         bucket.Bucket,
			}, object, nil)
			go s.matchSubscriptions(&protos.SubscriptionEvent{
				SubscriptionType: protos.SubscriptionType_BUCKET,
				BucketID:         bucket.Bucket,
			}, nil, bucket)
		}
	}
}

func (s *Service) Subscribe(subscription *protos.SubscriptionEvent,
	stream protos.MetadataService_SubscribeServer) error {
	logger.InfoLogger.Println("got subscription %+v ", subscription)
	var err error
	finished := make(chan bool)
	if subscription.SubscriptionType == protos.SubscriptionType_BUCKET {
		s.bucketSubscribersLock.Lock()
		s.bucketSubscribers[subscription.BucketID] = append(s.bucketSubscribers[subscription.BucketID], subscription.SubscriberID)
		s.bucketSubscriberStreams[subscription.SubscriberID] = &SubscriberStream{
			stream:   stream,
			finished: finished,
		}
		s.bucketSubscribersLock.Unlock()
	} else if subscription.SubscriptionType == protos.SubscriptionType_OBJECT {
		s.objectSubscribersLock.Lock()
		objectId := s.createObjectKey(subscription)
		s.objectSubscribers[objectId] = append(s.objectSubscribers[objectId], subscription.SubscriberID)
		s.objectSubscriberStreams[subscription.SubscriberID] = &SubscriberStream{
			stream:   stream,
			finished: finished,
		}
		s.objectSubscribersLock.Unlock()
	} else if subscription.SubscriptionType == protos.SubscriptionType_PREFIX {
		s.prefixSubscribersLock.Lock()
		objectId := s.createObjectKey(subscription)
		s.prefixSubscribers[objectId] = append(s.prefixSubscribers[objectId], subscription.SubscriberID)
		s.prefixSubscriberStreams[subscription.SubscriberID] = &SubscriberStream{
			stream:   stream,
			finished: finished,
		}
		s.prefixSubscribersLock.Unlock()
	} else {
		err = errors.New("subscription type not found")
		logger.ErrorLogger.Println(err)
		return err
	}
	cntx := stream.Context()
	for {
		select {
		case <-finished:
			return nil
		case <-cntx.Done():
			return nil
		}
	}
}

func (s *Service) matchSubscriptions(subscription *protos.SubscriptionEvent,
	object *protos.Object, bucket *protos.Bucket) {
	var subscribers []string
	var currentSubscribers []string
	var ok bool
	if subscription.SubscriptionType == protos.SubscriptionType_BUCKET {
		s.bucketSubscribersLock.RLock()
		if currentSubscribers, ok = s.bucketSubscribers[subscription.BucketID]; ok {
			subscribers = append(subscribers, currentSubscribers...)
		}
		s.bucketSubscribersLock.RUnlock()
	} else if subscription.SubscriptionType == protos.SubscriptionType_OBJECT {
		s.objectSubscribersLock.RLock()
		objectId := s.createObjectKey(subscription)
		if currentSubscribers, ok = s.objectSubscribers[objectId]; ok {
			subscribers = append(subscribers, currentSubscribers...)
		}
		s.objectSubscribersLock.RUnlock()
	} else if subscription.SubscriptionType == protos.SubscriptionType_PREFIX {
		s.prefixSubscribersLock.RLock()
		objectId := s.createObjectKey(subscription)
		if currentSubscribers, ok = s.prefixSubscribers[objectId]; ok {
			subscribers = append(subscribers, currentSubscribers...)
		}
		s.prefixSubscribersLock.RUnlock()
	}
	if !ok {
		return
	}
	for _, subscriberID := range subscribers {
		s.sendSubscriptions(subscription, subscriberID, object, bucket)
	}
}

func (s *Service) sendSubscriptions(subscription *protos.SubscriptionEvent, subscriberID string,
	object *protos.Object, bucket *protos.Bucket) {
	var streamer *SubscriberStream
	var ok bool
	if subscription.SubscriptionType == protos.SubscriptionType_BUCKET {
		s.bucketSubscribersLock.RLock()
		streamer, ok = s.bucketSubscriberStreams[subscriberID]
		s.bucketSubscribersLock.RUnlock()
	} else if subscription.SubscriptionType == protos.SubscriptionType_OBJECT {
		s.objectSubscribersLock.RLock()
		streamer, ok = s.objectSubscriberStreams[subscriberID]
		s.objectSubscribersLock.RUnlock()
	} else if subscription.SubscriptionType == protos.SubscriptionType_PREFIX {
		s.prefixSubscribersLock.RLock()
		streamer, ok = s.prefixSubscriberStreams[subscriberID]
		s.prefixSubscribersLock.RUnlock()
	}
	if !ok {
		return
	}
	if err := streamer.stream.Send(&protos.SubscriptionResponse{
		SubscriptionType: subscription.SubscriptionType,
		Bucket:           bucket,
		Object:           object,
	}); err != nil {
		logger.ErrorLogger.Println("could not send the proposal response to subscriber " + subscriberID)
		s.removeSubscriber(streamer, subscription, subscriberID)
	}
	logger.InfoLogger.Println("sending subscription: ", subscription.BucketID,
		subscription.Key, subscriberID)
}

func (s *Service) removeSubscriber(streamer *SubscriberStream, subscription *protos.SubscriptionEvent, subscriberID string) {
	streamer.finished <- true
	if subscription.SubscriptionType == protos.SubscriptionType_BUCKET {
		s.bucketSubscribersLock.Lock()
		delete(s.bucketSubscriberStreams, subscription.BucketID)
		if currentSubscribers, ok := s.bucketSubscribers[subscription.BucketID]; ok {
			s.removeElementFromSlice(currentSubscribers, subscriberID)
		}
		s.bucketSubscribersLock.Unlock()
	} else if subscription.SubscriptionType == protos.SubscriptionType_OBJECT {
		s.objectSubscribersLock.Lock()
		delete(s.objectSubscriberStreams, subscription.Key)
		objectId := s.createObjectKey(subscription)
		if currentSubscribers, ok := s.objectSubscribers[objectId]; ok {
			s.removeElementFromSlice(currentSubscribers, subscriberID)
		}
		s.objectSubscribersLock.Unlock()
	} else if subscription.SubscriptionType == protos.SubscriptionType_PREFIX {
		s.prefixSubscribersLock.Lock()
		objectId := s.createObjectKey(subscription)
		delete(s.prefixSubscriberStreams, subscription.Key+subscription.Prefix)
		if currentSubscribers, ok := s.prefixSubscribers[objectId]; ok {
			s.removeElementFromSlice(currentSubscribers, subscriberID)
		}
		s.prefixSubscribersLock.Unlock()
	}
}

func (s *Service) removeElementFromSlice(subscribers []string, subscriberID string) {
	var index int
	for subscriberIndex, currentSubscriberID := range subscribers {
		if subscriberID == currentSubscriberID {
			index = subscriberIndex
			break
		}
	}
	subscribers[index] = subscribers[len(subscribers)-1]
	subscribers[len(subscribers)-1] = ""
	subscribers = subscribers[:len(subscribers)-1]
}

func (s *Service) Unsubscribe(unsubscription *protos.SubscriptionEvent) error {
	logger.InfoLogger.Println("got unsubscribe %+v ", unsubscription)
	var streamer *SubscriberStream
	var ok bool
	if unsubscription.SubscriptionType == protos.SubscriptionType_BUCKET {
		s.bucketSubscribersLock.RLock()
		streamer, ok = s.bucketSubscriberStreams[unsubscription.SubscriberID]
		s.bucketSubscribersLock.RUnlock()
	} else if unsubscription.SubscriptionType == protos.SubscriptionType_OBJECT {
		s.objectSubscribersLock.RLock()
		streamer, ok = s.objectSubscriberStreams[unsubscription.SubscriberID]
		s.objectSubscribersLock.RUnlock()
	} else if unsubscription.SubscriptionType == protos.SubscriptionType_PREFIX {
		s.prefixSubscribersLock.RLock()
		streamer, ok = s.prefixSubscriberStreams[unsubscription.SubscriberID]
		s.prefixSubscribersLock.RUnlock()
	}
	if !ok {
		return nil
	}
	s.removeSubscriber(streamer, unsubscription, unsubscription.SubscriberID)
	return nil
}
