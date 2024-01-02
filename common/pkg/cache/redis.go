package commoncache

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func NewRedis(ctx context.Context, addr string) (*redis.Client, error) {
	// 创建 Redis 客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	// 检查是否有错误
	if err := redisClient.Ping(ctx).Err(); err != nil {
		// 返回错误信息
		return nil, err
	}

	return redisClient, nil
}
