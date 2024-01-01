package cache

import "github.com/go-redis/redis/v8"

func NewRedis(addr string) *redis.Client {
	// 创建 Redis 客户端
	redisClinet := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	return redisClinet
}
