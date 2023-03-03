package main

import (
	"context"
	"fmt"
	"time"

	c "github.com/lzzzzl/gogo-learn/redis"
	"github.com/redis/go-redis/v9"
)

var (
	ctx    = context.Background()
	client = c.Client1()
)

func pipe1() {
	// Connect to Redis Server
	pipe := client.Pipeline()

	incr := pipe.Incr(ctx, "pipeline_counter")
	pipe.Expire(ctx, "pipeline_counter", time.Hour)

	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}

	// The value is available only after Exec is called.
	fmt.Println(incr.Val())
}

func pipe2() {
	var incr *redis.IntCmd

	_, err := client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		incr = pipe.Incr(ctx, "pipeline_counter")
		pipe.Expire(ctx, "pipeline_counter", time.Hour)
		return nil
	})
	if err != nil {
		panic(err)
	}

	// The value is available only after the pipeline is executed.
	fmt.Println(incr.Val())
}

// Pipelines also return the executed commands so can iterate over them to retrieve results.
func pipe3() {
	cmds, err := client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for i := 0; i < 100; i++ {
			pipe.Get(ctx, fmt.Sprintf("key%d", i))
		}
		return nil
	})
	if err != nil {
		fmt.Println("pipe3 err: ", err)
	}

	for _, cmd := range cmds {
		fmt.Println(cmd.(*redis.StringCmd).Val())
	}
}

func main() {
	pipe1()
	pipe2()
	pipe3()

	client.Close()
}
