package database

/*
	temporary store to hold tokens
*/

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewClient(ctx context.Context, address string) *redis.Client {
	log.Printf("client starting")
	client := redis.NewClient(&redis.Options{Addr: address})
	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Printf("client initialization failed %s", err)
	}
	log.Printf("Connected to store: %s", address)
	return client
}
