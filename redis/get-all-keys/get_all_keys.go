package main

import (
	"context"
	"fmt"

	c "github.com/lzzzzl/gogo-learn/redis"
)

var (
	ctx    = context.Background()
	client = c.Client1()
)

func scan1() {
	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = client.Scan(ctx, cursor, "prefix:*", 0).Result()
		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			fmt.Println("key: ", key)
		}

		// no more keys
		if cursor == 0 {
			break
		}
	}
}

// simplify the code
func scan2() {
	iter := client.Scan(ctx, 0, "prefix:*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys: ", iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}

// Sets and hashes
func scan3() {
	iter := client.SScan(ctx, "set-key", 0, "prefix:*", 0).Iterator()
	iter = client.HScan(ctx, "hash-key", 0, "prefix:*", 0).Iterator()
	iter = client.ZScan(ctx, "sorted-hash-key", 0, "prefix:*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys: ", iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}

// Delete keys without TTL
func scan4() {
	iter := client.Scan(ctx, 0, "prefix:*", 0).Iterator()

	for iter.Next(ctx) {
		key := iter.Val()

		d, err := client.TTL(ctx, key).Result()
		if err != nil {
			panic(err)
		}

		// -1 means no TTL
		if d == -1 {
			if err := client.Del(ctx, key).Err(); err != nil {
				panic(err)
			}
		}
	}

	if err := iter.Err(); err != nil {
		panic(err)
	}
}

func main() {
	scan1()
	scan2()
	scan3()
}
