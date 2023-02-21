package mds

import (
	"context"
	"github.com/IBM/gedsmds/internal/connection/connpool"
	"github.com/IBM/gedsmds/internal/logger"
	"github.com/IBM/gedsmds/internal/mdsprocessor"
	"github.com/IBM/gedsmds/protos/protos"
)

type Service struct {
	mdsProcessor *mdsprocessor.Processor
}

func NewService() *Service {
	return &Service{
		mdsProcessor: mdsprocessor.InitProcessor(),
	}
}

func (t *Service) GetConnectionInformation(_ context.Context, _ *protos.EmptyParams) (*protos.ConnectionInformation, error) {
	currentIP := connpool.GetOutboundIP()
	logger.InfoLogger.Println("Found my IP:", currentIP)
	return &protos.ConnectionInformation{
		RemoteAddress: currentIP,
	}, nil
}

func (t *Service) RegisterObjectStore(_ context.Context, _ *protos.ObjectStoreConfig) (*protos.StatusResponse, error) {
	return nil, nil
}

func (t *Service) ListObjectStores(_ context.Context, _ *protos.EmptyParams) (*protos.AvailableObjectStoreConfigs, error) {
	return nil, nil
}

func (t *Service) CreateBucket(_ context.Context, _ *protos.Bucket) (*protos.StatusResponse, error) {
	return nil, nil
}

func (t *Service) DeleteBucket(_ context.Context, _ *protos.Bucket) (*protos.StatusResponse, error) {
	return nil, nil
}

func (t *Service) ListBuckets(_ context.Context, _ *protos.EmptyParams) (*protos.BucketListResponse, error) {
	return nil, nil
}

func (t *Service) LookupBucket(_ context.Context, _ *protos.Bucket) (*protos.StatusResponse, error) {
	return nil, nil
}

func (t *Service) Create(_ context.Context, object *protos.Object) (*protos.StatusResponse, error) {
	logger.InfoLogger.Println("%+v\n", object)
	return &protos.StatusResponse{
		Code: protos.StatusCode_OK,
	}, nil
}

func (t *Service) Update(_ context.Context, _ *protos.Object) (*protos.StatusResponse, error) {
	return nil, nil
}

func (t *Service) Delete(_ context.Context, _ *protos.ObjectID) (*protos.StatusResponse, error) {
	return nil, nil
}

func (t *Service) DeletePrefix(_ context.Context, _ *protos.ObjectID) (*protos.StatusResponse, error) {
	return nil, nil
}

func (t *Service) Lookup(_ context.Context, _ *protos.ObjectID) (*protos.ObjectResponse, error) {
	return nil, nil
}

func (t *Service) List(_ context.Context, _ *protos.ObjectListRequest) (*protos.ObjectListResponse, error) {
	return nil, nil
}
