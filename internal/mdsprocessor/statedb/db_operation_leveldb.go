package statedb

import (
	"github.com/IBM/gedsmds/internal/logger"
	"github.com/syndtr/goleveldb/leveldb"
)

const dbsLocation = "./data/state_"

type Operations struct {
	db *leveldb.DB
}

type ByteContainer struct {
	Value []byte
}

type ByteContainers struct {
	Values []*ByteContainer
}

// NewOperations Possible optimization for LevelDB: https://github.com/google/leveldb/blob/master/doc/index.md
func NewOperations(contractName string) *Operations {
	tempDB, err := leveldb.OpenFile(dbsLocation+contractName, nil)
	if err != nil {
		logger.FatalLogger.Fatalln(err)
	}
	return &Operations{
		db: tempDB,
	}
}
