package cache

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Cache struct {
	client *redis.Client
}

func Connect() *Cache {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		),
	})

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatal("redis not reachable:", err)
	}

	log.Println("Redis connected successfully")
	return &Cache{client: client}
}

func (c *Cache) Set(key, value string, ttl time.Duration) error {
	return c.client.Set(ctx, key, value, ttl).Err()
}

func (c *Cache) Get(key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *Cache) Delete(key string) error {
	return c.client.Del(ctx, key).Err()
}