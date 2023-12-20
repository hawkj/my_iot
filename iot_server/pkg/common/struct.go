package common

import (
	"github.com/go-redis/redis/v8"
	"github.com/hawkj/my_iot/iot_server/config"
)

type Global struct {
	Config *config.Config
	Redis  *redis.Client
}
