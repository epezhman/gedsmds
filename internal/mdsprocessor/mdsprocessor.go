package mdsprocessor

import (
	"github.com/IBM/gedsmds/internal/config"
	"github.com/IBM/gedsmds/internal/connection/connpool"
	"github.com/IBM/gedsmds/internal/keyvaluestore"
	"github.com/IBM/gedsmds/internal/logger"
	"github.com/IBM/gedsmds/internal/pubsub"
	"github.com/IBM/gedsmds/protos/protos"
)

func InitService() *Service {
	kvStore := keyvaluestore.InitKeyValueStoreService()
	return &Service{
		pubsub:  pubsub.InitService(kvStore),
		kvStore: kvStore,
	}
}

func (s *Service) GetConnectionInformation() string {
	currentIP := connpool.GetOutboundIP()
	logger.InfoLogger.Println("found my IP:", currentIP)
	return currentIP
}

func (s *Service) RegisterObjectStore(objectStore *protos.ObjectStoreConfig) error {
	if err := s.kvStore.RegisterObjectStore(objectStore); err != nil {
		return err
	}
	logger.InfoLogger.Println("objectStore create %+v", objectStore)
	return nil
}

func (s *Service) ListObjectStores() (*protos.AvailableObjectStoreConfigs, error) {
	return s.kvStore.ListObjectStores()
}

func (s *Service) CreateBucket(bucket *protos.Bucket) error {
	if err := s.kvStore.CreateBucket(bucket); err != nil {
		return err
	}
	logger.InfoLogger.Println("bucket create %+v", bucket)
	return nil
}

func (s *Service) DeleteBucket(bucket *protos.Bucket) error {
	if err := s.kvStore.DeleteBucket(bucket); err != nil {
		return err
	}
	logger.InfoLogger.Println("bucket delete %+v", bucket)
	return nil
}

func (s *Service) ListBuckets() (*protos.BucketListResponse, error) {
	return s.kvStore.ListBuckets()
}

func (s *Service) LookupBucket(bucket *protos.Bucket) error {
	if err := s.kvStore.LookupBucket(bucket); err != nil {
		return err
	}
	return nil
}

func (s *Service) CreateObject(object *protos.Object) error {
	if err := s.kvStore.CreateObject(object); err != nil {
		return err
	}
	logger.InfoLogger.Println("object create %+v", object)

	if config.Config.PubSubEnabled {
		s.pubsub.UpdatedObject <- object
	}
	return nil
}

func (s *Service) CreateObjectStream(object *protos.Object) {
	if err := s.kvStore.CreateObject(object); err != nil {
		return
	}
	logger.InfoLogger.Println("object create %+v", object)
	if config.Config.PubSubEnabled {
		s.pubsub.UpdatedObject <- object
	}
}

func (s *Service) UpdateObject(object *protos.Object) error {
	if err := s.kvStore.UpdateObject(object); err != nil {
		return err
	}
	if config.Config.PubSubEnabled {
		s.pubsub.UpdatedObject <- object
	}
	return nil
}

func (s *Service) UpdateObjectStream(object *protos.Object) {
	if err := s.kvStore.UpdateObject(object); err != nil {
		return
	}
	if config.Config.PubSubEnabled {
		s.pubsub.UpdatedObject <- object
	}
}

func (s *Service) DeleteObject(objectID *protos.ObjectID) error {
	if err := s.kvStore.DeleteObject(objectID); err != nil {
		return err
	}
	if config.Config.PubSubEnabled {
		s.pubsub.UpdatedObject <- &protos.Object{
			Id: objectID,
		}
	}
	return nil
}

func (s *Service) DeletePrefix(_ *protos.ObjectID) (*protos.StatusResponse, error) {
	return &protos.StatusResponse{Code: protos.StatusCode_UNIMPLEMENTED}, nil
}

func (s *Service) LookupObject(objectID *protos.ObjectID) (*protos.ObjectResponse, error) {
	object, err := s.kvStore.LookupObject(objectID)
	if err != nil {
		return &protos.ObjectResponse{
			Error: &protos.StatusResponse{Code: protos.StatusCode_NOT_FOUND},
		}, err
	}
	object.Error = &protos.StatusResponse{Code: protos.StatusCode_OK}
	return object, nil
}

func (s *Service) List(_ *protos.ObjectListRequest) (*protos.ObjectListResponse, error) {
	return &protos.ObjectListResponse{
		Error: &protos.StatusResponse{Code: protos.StatusCode_NOT_FOUND},
	}, nil
}

func (s *Service) Subscribe(subscription *protos.SubscriptionEvent,
	stream protos.MetadataService_SubscribeServer) error {
	return s.pubsub.Subscribe(subscription, stream)
}

func (s *Service) Unsubscribe(unsubscribe *protos.SubscriptionEvent) error {
	if err := s.pubsub.Unsubscribe(unsubscribe); err != nil {
		return err
	}
	return nil
}
