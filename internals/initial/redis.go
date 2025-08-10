package initial

import (
	"ecommerce_clean/pkgs/logger"
	"ecommerce_clean/pkgs/redis"
	"fmt"
)

func InitRedis(info redis.Config) (redis.IRedis, error) {
	redis := redis.NewRedis(info)
	if !redis.IsConnected() {
		logger.Fatal("Redis connection failed")
		return nil, fmt.Errorf("redis connection failed")
	}

	logger.Info("Redis connection established successfully")
	return redis, nil
}
