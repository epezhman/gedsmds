package db

import (
	"github.com/IBM/gedsmds/internal/logger"
	"github.com/IBM/gedsmds/protos/protos"
	"github.com/golang/protobuf/proto"
	"github.com/syndtr/goleveldb/leveldb"
)

const dbsLocation = "./data/state_mds"

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
func NewOperations() *Operations {
	tempDB, err := leveldb.OpenFile(dbsLocation, nil)
	if err != nil {
		logger.FatalLogger.Fatalln(err)
	}
	return &Operations{
		db: tempDB,
	}
}

func (o *Operations) PutObject(object *protos.Object) {
	dbValue, _ := proto.Marshal(object)
	if err := o.db.Put([]byte(object.Id.Bucket+object.Id.Key), dbValue, nil); err != nil {
		logger.ErrorLogger.Println(err)
	}
}
