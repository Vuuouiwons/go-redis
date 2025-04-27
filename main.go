package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-cache:6379",
		Password: "password", // No password set
		DB:       0,          // Use default DB
		Protocol: 2,          // Connection protocol
	})

	ctx := context.Background()
	
	// add / update key: test, with value: payload
	err := client.Set(ctx, "test", "payload", 0).Err()
	if err != nil {
		panic(err)
	}

	// get key: test
	val, err := client.Get(ctx, "test").Result()
	if err != nil {
		panic(err)
	}
	// print it to stdout
	fmt.Println("test", val)

	// remove key: test
	err = client.Del(ctx, "test").Err()
	if err != nil {
		panic(err)
	}

}
