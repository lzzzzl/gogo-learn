package main

import (
	"context"
	"fmt"

	c "github.com/lzzzzl/gogo-learn/redis"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	// Connect to Redis Server
	client := c.Client1()

	// Ping Redis to test the connection
	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)

	// Set a key-value pair
	err = client.Set(ctx, "myKey", "myValue", 0).Err()
	if err != nil {
		panic(err)
	}

	// Get the vaule of a key
	value, err := client.Get(ctx, "myKey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("myKey", value)

	// Delete a key
	err = client.Del(ctx, "myKey").Err()
	if err != nil {
		panic(err)
	}

	// executing unsupported commands
	val, err := client.Do(ctx, "get", "key").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exists")
		} else {
			panic(err)
		}
	} else {
		fmt.Println(val.(string))
	}

	// Close the connection
	err = client.Close()
	if err != nil {
		panic(err)
	}
}
