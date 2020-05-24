package athena

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

// RedisEnv is the definition of environment variables needed
// to connect to redis
type RedisEnv struct {
	RedisHost     string `required:"true" split_words:"true"`
	RedisPort     int    `required:"true" split_words:"true"`
	RedisPassword string `default:"" split_words:"true"`
}

// NewRedisClient creates a client for redis and tests its connection
func NewRedisClient(env RedisEnv) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", env.RedisHost, env.RedisPort),
		Password: env.RedisPassword,
		DB:       0,
	})

	// test the connection
	_, err := client.Ping().Result()

	return client, err
}
