package grpc


import (
  "fmt"
  "net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

  "multiplica/cmd/config"

)

func NewGRPCHandler(configuration config.Configurations) {

  l, err := net.Listen("tcp", fmt.Sprintf(":%d", configuration.Server.GrpcPort))
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("Failed to listen")
	}

  grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	RegisterMultiplyServiceServer(grpcServer, &MultiplyGrpcServer{})

  log.Info("gRPC server started at ", configuration.Server.GrpcPort)
	if err := grpcServer.Serve(l); err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("Failed to serve")
	}

}

//grpcurl -d '{"numberA": 2, "numberB": 12}' -plaintext localhost:9090 grpc.MultiplyService/Multiply 
