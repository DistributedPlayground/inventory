package store

import (
	"github.com/DistributedPlayground/go-lib/common"
	"github.com/DistributedPlayground/inventory/config"
	"github.com/DistributedPlayground/inventory/database"
)

func NewRedis() database.RedisStore {
	opts := database.RedisConfigOptions{
		Host:        config.Var.REDIS_HOST,
		Port:        config.Var.REDIS_PORT,
		Password:    config.Var.REDIS_PASSWORD,
		ClusterMode: !common.IsLocalEnv(),
	}
	return database.NewRedisStore(opts)
}
