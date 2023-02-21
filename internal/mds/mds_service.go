package mds

import (
	"context"
	"github.com/IBM/gedsmds/internal/logger"
	"github.com/IBM/gedsmds/internal/mdsprocessor"
	"github.com/IBM/gedsmds/protos/goprotos"
)

type Service struct {
	mdsProcessor *mdsprocessor.Processor
}

func NewService() *Service {
	return &Service{
		mdsProcessor: mdsprocessor.InitProcessor(),
	}
}

func (t *Service) SubscribeBucket(_ context.Context, bucket *protos.BucketEventSubscription) (*protos.BucketEventSubscription, error) {
	logger.InfoLogger.Println(bucket)
	return &protos.BucketEventSubscription{
		BucketId: bucket.BucketId,
	}, nil
}
