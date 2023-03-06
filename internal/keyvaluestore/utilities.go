package keyvaluestore

import (
	"github.com/IBM/gedsmds/protos/protos"
)

func (kv *Service) createObjectKey(object *protos.Object) string {
	return object.Id.Bucket + "/" + object.Id.Key
}
