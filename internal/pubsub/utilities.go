package pubsub

import (
	"github.com/IBM/gedsmds/protos/protos"
)

func (s *Service) createObjectKey(object *protos.SubscriptionEvent) string {
	return object.BucketID + "-" + object.Key
}
