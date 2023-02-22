package mdsservice

import (
	"context"
	"github.com/IBM/gedsmds/internal/connection/connpool"
	"github.com/IBM/gedsmds/internal/keyvaluestore"
	"github.com/IBM/gedsmds/internal/logger"
	"github.com/IBM/gedsmds/internal/mdsprocessor"
	"github.com/IBM/gedsmds/protos/protos"
)

func NewService() *Service {
	kvStore := keyvaluestore.InitKeyValueStore()
	return &Service{
		mdsProcessor: mdsprocessor.InitProcessor(kvStore),
		kvStore:      kvStore,
	}
}

func (s *Service) GetConnectionInformation(_ context.Context,
	_ *protos.EmptyParams) (*protos.ConnectionInformation, error) {
	currentIP := connpool.GetOutboundIP()
	logger.InfoLogger.Println("Found my IP:", currentIP)
	return &protos.ConnectionInformation{RemoteAddress: currentIP}, nil
}

func (s *Service) RegisterObjectStore(_ context.Context,
	objectStore *protos.ObjectStoreConfig) (*protos.StatusResponse, error) {
	if err := s.kvStore.RegisterObjectStore(objectStore); err != nil {
		return &protos.StatusResponse{Code: protos.StatusCode_ALREADY_EXISTS}, err
	}
	return &protos.StatusResponse{Code: protos.StatusCode_OK}, nil
}

func (s *Service) ListObjectStores(_ context.Context,
	_ *protos.EmptyParams) (*protos.AvailableObjectStoreConfigs, error) {
	return s.kvStore.ListObjectStores()
}

func (s *Service) CreateBucket(_ context.Context, bucket *protos.Bucket) (*protos.StatusResponse, error) {
	if err := s.kvStore.CreateBucket(bucket); err != nil {
		return &protos.StatusResponse{Code: protos.StatusCode_ALREADY_EXISTS}, err
	}
	return &protos.StatusResponse{Code: protos.StatusCode_OK}, nil
}

func (s *Service) DeleteBucket(_ context.Context, bucket *protos.Bucket) (*protos.StatusResponse, error) {
	if err := s.kvStore.DeleteBucket(bucket); err != nil {
		return &protos.StatusResponse{Code: protos.StatusCode_NOT_FOUND}, err
	}
	return &protos.StatusResponse{Code: protos.StatusCode_OK}, nil
}

func (s *Service) ListBuckets(_ context.Context, _ *protos.EmptyParams) (*protos.BucketListResponse, error) {
	return s.kvStore.ListBuckets()
}

func (s *Service) LookupBucket(_ context.Context, bucket *protos.Bucket) (*protos.StatusResponse, error) {
	if err := s.kvStore.LookupBucket(bucket); err != nil {
		return &protos.StatusResponse{Code: protos.StatusCode_NOT_FOUND}, err
	}
	return &protos.StatusResponse{Code: protos.StatusCode_OK}, nil
}

func (s *Service) Create(_ context.Context, object *protos.Object) (*protos.StatusResponse, error) {
	if err := s.kvStore.CreateObject(object); err != nil {
		return &protos.StatusResponse{Code: protos.StatusCode_ALREADY_EXISTS}, err
	}
	return &protos.StatusResponse{Code: protos.StatusCode_OK}, nil
}

func (s *Service) Update(_ context.Context, object *protos.Object) (*protos.StatusResponse, error) {
	if err := s.kvStore.UpdateObject(object); err != nil {
		return &protos.StatusResponse{Code: protos.StatusCode_ALREADY_EXISTS}, err
	}
	return &protos.StatusResponse{Code: protos.StatusCode_OK}, nil
}

func (s *Service) Delete(_ context.Context, objectID *protos.ObjectID) (*protos.StatusResponse, error) {
	if err := s.kvStore.DeleteObject(objectID); err != nil {
		return &protos.StatusResponse{Code: protos.StatusCode_NOT_FOUND}, err
	}
	return &protos.StatusResponse{Code: protos.StatusCode_OK}, nil
}

func (s *Service) DeletePrefix(_ context.Context, _ *protos.ObjectID) (*protos.StatusResponse, error) {
	return &protos.StatusResponse{Code: protos.StatusCode_UNIMPLEMENTED}, nil
}

func (s *Service) Lookup(_ context.Context, objectID *protos.ObjectID) (*protos.ObjectResponse, error) {
	object, err := s.kvStore.LookupObject(objectID)
	if err != nil {
		return &protos.ObjectResponse{
			Error: &protos.StatusResponse{Code: protos.StatusCode_NOT_FOUND},
		}, err
	}
	object.Error = &protos.StatusResponse{Code: protos.StatusCode_OK}
	return object, nil
}

func (s *Service) List(_ context.Context, _ *protos.ObjectListRequest) (*protos.ObjectListResponse, error) {
	return &protos.ObjectListResponse{
		Error: &protos.StatusResponse{Code: protos.StatusCode_NOT_FOUND},
	}, nil
}

func (s *Service) TestRPC(_ context.Context, conn *protos.ConnectionInformation) (*protos.ConnectionInformation, error) {
	logger.InfoLogger.Println("Got this:", conn.RemoteAddress)
	return &protos.ConnectionInformation{RemoteAddress: conn.RemoteAddress}, nil
}

func (s *Service) SubscribeObjects(subscription *protos.ObjectEventSubscription,
	stream protos.MetadataService_SubscribeObjectsServer) error {
	return s.mdsProcessor.ObjectEventSubscription(subscription, stream)
}
