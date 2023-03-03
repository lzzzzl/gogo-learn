package main

import (
	"context"
	"fmt"

	c "github.com/lzzzzl/gogo-learn/redis"
	"github.com/redis/go-redis/v9"
)

type Model struct {
	Str1   string   `redis:"str1"`
	Str2   string   `redis:"str2"`
	Int    int      `redis:"int"`
	Bool   bool     `redis:"bool"`
	Ignore struct{} `redis:"-"`
}

var (
	ctx    = context.Background()
	client = c.Client1()
)

func prepare() {
	if _, err := client.Pipelined(ctx, func(p redis.Pipeliner) error {
		client.HSet(ctx, "key", "str1", "hello")
		client.HSet(ctx, "key", "str2", "world")
		client.HSet(ctx, "key", "int", 123)
		client.HSet(ctx, "key", "bool", 1)
		return nil
	}); err != nil {
		panic(err)
	}
}

func scan1() {
	var model1 Model
	// Scan all fields into the model
	if err := client.HGetAll(ctx, "key").Scan(&model1); err != nil {
		panic(err)
	}
	fmt.Println(model1)
}

func scan2() {
	var model2 Model
	// Scan a subset of the fields
	if err := client.HMGet(ctx, "key", "str1", "str2").Scan(&model2); err != nil {
		panic(err)
	}
	fmt.Println(model2)
}

func main() {
	prepare()
	scan1()
	scan2()
}
