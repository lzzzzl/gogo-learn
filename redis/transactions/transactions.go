package main

import (
	"context"
	"errors"
	"fmt"

	c "github.com/lzzzzl/gogo-learn/redis"
	"github.com/redis/go-redis/v9"
)

var (
	ctx    = context.Background()
	client = c.Client1()
)

// You can wrap a pipeline with MULTI and EXEC commands using TxPipelined and TxPipeline,
// but it is not very useful on its own:
func transactions1() {
	_, err := client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		for i := 0; i < 100; i++ {
			pipe.Get(ctx, fmt.Sprintf("key%d", i))
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

// Instead, you should transactional pipelines with Watch, for example,
// we can correctly implement INCR command using GET, SET, and WATCH.
// Note how we use redis.TxFailedErr to check if the transaction has failed or not.
func transactions2(key string) error {
	// Redis transactions use optimistic locking.
	const maxRetries = 1000

	// Increment transactionally increments the key using GET and SET commands.
	// Transactional function.
	txf := func(tx *redis.Tx) error {
		// Get the current value or zero.
		n, err := tx.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			return err
		}

		// Actual operation (local in optimistic lock).
		n++

		// Operation is commited only if the watched keys remain unchanged.
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, key, n, 0)
			return nil
		})
		return err
	}

	// Retry if the key has been changed.
	for i := 0; i < maxRetries; i++ {
		err := client.Watch(ctx, txf, key)
		if err == nil {
			// Success.
			return nil
		}
		if err == redis.TxFailedErr {
			// Optimistic lock lost. Retry.
			continue
		}
		// Return any other error.
		return err
	}

	return errors.New("increment reached maximum number of retries")
}

func main() {
	transactions1()

	if err := transactions2("test"); err != nil {
		fmt.Println(err)
	}

	client.Close()
}
