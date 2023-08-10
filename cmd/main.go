package main

import (
	"log"
	"os"

	env "github.com/DistributedPlayground/go-lib/config"
	"github.com/DistributedPlayground/inventory/api"
	"github.com/DistributedPlayground/inventory/config"
	"github.com/DistributedPlayground/inventory/pkg/service"
	"github.com/DistributedPlayground/inventory/pkg/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	// load env vars
	err := env.LoadEnv(&config.Var)
	if err != nil {
		log.Fatalf("failed to load env vars: %v", err)
	}
	lg := zerolog.New(os.Stdout)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	port := config.Var.PORT
	store.NewRedis()
	api.Start(api.APIConfig{
		Port:    port,
		Logger:  &lg,
		Service: service.NewInventory(redis),
	})
}
