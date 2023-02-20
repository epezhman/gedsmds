package main

import (
	"github.com/IBM/gedsmds/internal/config"
	"github.com/IBM/gedsmds/internal/connection/serverconfig"
	"github.com/IBM/gedsmds/internal/logger"
	"github.com/IBM/gedsmds/internal/mds"
	protos "github.com/IBM/gedsmds/protos/goprotos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	_ "google.golang.org/grpc/encoding/gzip"
	"net"
	"path/filepath"
)

func main() {
	cert, _ := filepath.Abs("./configs/cert.pem")
	key, _ := filepath.Abs("./configs/key.pem")
	lis, err := net.Listen("tcp", config.Config.MDSServerPort)
	if err != nil {
		logger.FatalLogger.Fatalln(err)
	}
	credential, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		logger.FatalLogger.Fatalln(err)
	}
	opts := []grpc.ServerOption{grpc.KeepaliveEnforcementPolicy(serverconfig.KAEP),
		grpc.KeepaliveParams(serverconfig.KASP), grpc.Creds(credential)}
	grpcServer := grpc.NewServer(opts...)
	serviceInstance := mds.NewService()
	protos.RegisterMDSServiceServer(grpcServer, serviceInstance)
	logger.InfoLogger.Println("Transaction Server is listening on port", config.Config.MDSServerPort)
	err = grpcServer.Serve(lis)
	if err != nil {
		logger.FatalLogger.Fatalln(err)
	}
}
