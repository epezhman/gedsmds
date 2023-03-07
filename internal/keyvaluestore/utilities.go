package keyvaluestore

import (
	"github.com/IBM/gedsmds/protos/protos"
	"strings"
)

//func (kv *Service) createObjectKey(objectId *protos.ObjectID) string {
//	return objectId.Bucket + "-" + objectId.Key
//}

func (kv *Service) getNestedPath(objectId *protos.ObjectID) []string {
	return strings.Split(objectId.GetKey(), commonDelimiter)
}
