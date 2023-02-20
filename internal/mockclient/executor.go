package mockclient

import (
	"context"
	"github.com/IBM/gedsmds/internal/connection/connpool"
	"github.com/IBM/gedsmds/internal/logger"
	protos "github.com/IBM/gedsmds/protos/goprotos"
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
	conn, err := e.mdsConnections["localhost"].Get(context.Background())
	if conn == nil || err != nil {
		logger.ErrorLogger.Println(err)
	}
	client := protos.NewMDSServiceClient(conn.ClientConn)
	result, err := client.SubscribeBucket(context.Background(), &protos.BucketEventSubscription{BucketId: "test"})
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	logger.InfoLogger.Println(result)
	if errCon := conn.Close(); errCon != nil {
		logger.ErrorLogger.Println(errCon)
	}
}
