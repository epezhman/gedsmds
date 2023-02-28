package pubsub

import (
	"github.com/IBM/gedsmds/protos/protos"
)

func (s *Service) createObjectKey(object *protos.SubscriptionEvent) string {
	return object.BucketID + "-" + object.Key
}

//func (s *Service) createSubscriptionKey(subscription *protos.SubscriptionEvent) (string, error) {
//	var subscriptionID string
//	if subscription.SubscriptionType == protos.SubscriptionType_BUCKET {
//		subscriptionID = subscription.BucketID + "-" + subscription.SubscriberID
//	} else if subscription.SubscriptionType == protos.SubscriptionType_OBJECT {
//		subscriptionID = subscription.BucketID + "-" + subscription.Key + "-" + subscription.SubscriberID
//	} else if subscription.SubscriptionType == protos.SubscriptionType_PREFIX {
//		subscriptionID = subscription.BucketID + "-" + subscription.Key +
//			"-" + subscription.Prefix + "-" + subscription.SubscriberID
//	} else {
//		return subscriptionID, errors.New("subscription type not found")
//	}
//	return subscriptionID, nil
//}
