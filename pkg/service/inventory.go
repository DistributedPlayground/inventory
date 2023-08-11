package service

import (
	"context"

	"github.com/DistributedPlayground/inventory/api"
	"github.com/DistributedPlayground/inventory/database"
)

type inventory struct {
	redis database.RedisStore
	api.UnimplementedInventoryServer
}

func NewInventory(redis database.RedisStore, u api.UnimplementedInventoryServer) api.InventoryServer {
	return &inventory{redis, u}
}

func (i *inventory) Get(ctx context.Context, request *api.InventoryRequest) (response *api.InventoryResponse, err error) {
	count, err := i.redis.Get(request.Id)
	if err != nil {
		return &api.InventoryResponse{}, err
	} else {
		return &api.InventoryResponse{
			Count: string(count),
		}, nil
	}
}
