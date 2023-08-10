package api

import (
	"log"
	"net"

	"github.com/DistributedPlayground/inventory/pkg/service"
	"github.com/rs/zerolog"
	grpc "google.golang.org/grpc"
)

type APIConfig struct {
	Logger  *zerolog.Logger
	Port    string
	Service service.Inventory
}

func Start(config APIConfig) {
	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Create a new gRPC server
	RegisterInventoryServer(s, config.Service)
}
