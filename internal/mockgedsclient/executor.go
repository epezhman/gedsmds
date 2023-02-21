package mockgedsclient

import (
	"context"
	"github.com/IBM/gedsmds/internal/connection/connpool"
	"github.com/IBM/gedsmds/internal/logger"
	"github.com/IBM/gedsmds/protos/protos"
)

type Executor struct {
	mdsConnections map[string]*connpool.Pool
}

func NewExecutor() *Executor {
	tempEx := &Executor{
		mdsConnections: connpool.GetMDSConnectionsStream(),
	}
	connpool.SleepAndContinue()
	return tempEx
}

func (e *Executor) SendSubscription() {
	conn, err := e.mdsConnections["127.0.0.1"].Get(context.Background())
	if conn == nil || err != nil {
		logger.ErrorLogger.Println(err)
	}
	client := protos.NewMetadataServiceClient(conn.ClientConn)
	result, err := client.Create(context.Background(), &protos.Object{Id: &protos.ObjectID{Key: "test"}})
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	logger.InfoLogger.Println(result.Code)
	if errCon := conn.Close(); errCon != nil {
		logger.ErrorLogger.Println(errCon)
	}
}
