package databases

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func GetRedisClient(host string, port string) *redis.Client {
	opt, err := redis.ParseURL(fmt.Sprintf("%s:%s", host, port)) // "redis://redis:6379"
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opt)
}
