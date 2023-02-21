package connpool

import (
	"github.com/IBM/gedsmds/internal/config"
	"github.com/IBM/gedsmds/internal/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

var KACP = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             10 * time.Second, // wait 30 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

func factoryNode(ip string) (*grpc.ClientConn, error) {
	//configTLS := &tls.Config{
	//	InsecureSkipVerify: false,
	//	RootCAs:            certificates.CAs,
	//}
	//opts := []grpc.DialOption{grpc.WithTransportCredentials(credentials.NewTLS(configTLS)), grpc.WithKeepaliveParams(KACP),
	//	grpc.WithDefaultCallOptions(grpc.UseCompressor("gzip")),
	//	// https://chromium.googlesource.com/external/github.com/grpc/grpc-go/+/HEAD/Documentation/encoding.md
	//}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(ip+config.Config.MDSPort, opts...)
	if err != nil {
		logger.FatalLogger.Fatalln("Failed to start gRPC connection:", err)
	}
	return conn, err
}

func GetMDSConnectionsStream() map[string]*Pool {
	serverPool := make(map[string]*Pool)
	name := "127.0.0.1"
	pool, err := NewPoolWithIP(factoryNode, name, 1, 1, 10*time.Second)
	if err != nil {
		logger.FatalLogger.Fatalln("Failed to create gRPC pool:", err)
	}
	serverPool[name] = pool
	return serverPool
}

func SleepAndContinue() {
	time.Sleep(2 * time.Second)
}

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	defer func(conn net.Conn) {
		if errCon := conn.Close(); errCon != nil {
			logger.ErrorLogger.Println(errCon)
		}
	}(conn)
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
