package mdsservice

import (
	"github.com/IBM/gedsmds/internal/keyvaluestore"
	"github.com/IBM/gedsmds/internal/mdsprocessor"
)

type Service struct {
	mdsProcessor *mdsprocessor.Processor
	kvStore      *keyvaluestore.KeyValueStore
}
