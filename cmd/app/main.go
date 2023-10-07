package main

import (
	"log"
	"net"

	env "github.com/DistributedPlayground/go-lib/config"
	"github.com/DistributedPlayground/inventory/api"
	"github.com/DistributedPlayground/inventory/config"
	"github.com/DistributedPlayground/inventory/pkg/message"
	"github.com/DistributedPlayground/inventory/pkg/service"
	"github.com/DistributedPlayground/inventory/pkg/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// load env vars
	err := env.LoadEnv(&config.Var)
	if err != nil {
		log.Fatalf("failed to load env vars: %v", err)
	}
	// lg := zerolog.New(os.Stdout)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	listener, err := net.Listen("tcp", ":"+config.Var.PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	redis := store.NewRedis()

	messages := message.NewMessage(store.MustNewKafka(), redis)
	go messages.Listen()

	// Create a new instance of the inventory service
	service := service.NewInventory(redis, api.UnimplementedInventoryServer{})

	server := grpc.NewServer()

	// Register reflection service on the server.
	reflection.Register(server)

	// Create a new gRPC server
	api.RegisterInventoryServer(server, service)
	server.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
