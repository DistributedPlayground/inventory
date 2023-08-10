package service

import (
	"context"

	"github.com/DistributedPlayground/inventory/database"
	"github.com/DistributedPlayground/inventory/pkg/model"
)

type Inventory interface {
	Get(ctx context.Context, request *model.InventoryRequest) (response *model.InventoryResponse, err error)
}

type inventory struct {
	redis database.RedisStore
}

func NewInventory(redis database.RedisStore) Inventory {
	return &inventory{redis}
}

func (i *inventory) Get(ctx context.Context, request *model.InventoryRequest) (response *model.InventoryResponse, err error) {
	count, err := i.redis.Get(request.Id)
	if err != nil {
		return &model.InventoryResponse{}, err
	} else {
		return &model.InventoryResponse{
			Count: string(count),
		}, nil
	}
}
