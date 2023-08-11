package main

import (
	"log"
	"net"

	env "github.com/DistributedPlayground/go-lib/config"
	"github.com/DistributedPlayground/inventory/api"
	"github.com/DistributedPlayground/inventory/config"
	"github.com/DistributedPlayground/inventory/pkg/service"
	"github.com/DistributedPlayground/inventory/pkg/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"google.golang.org/grpc"
)

func main() {
	// load env vars
	err := env.LoadEnv(&config.Var)
	if err != nil {
		log.Fatalf("failed to load env vars: %v", err)
	}
	// lg := zerolog.New(os.Stdout)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	listener, err := net.Listen("tcp", config.Var.PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new instance of the inventory service
	service := service.NewInventory(store.NewRedis(), api.UnimplementedInventoryServer{})

	serverRegistrar := grpc.NewServer()

	// Create a new gRPC server
	api.RegisterInventoryServer(serverRegistrar, service)
	serverRegistrar.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
