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

func main() {
	_ = client.FlushDB(ctx).Err()

	for i := 0; i < 10; i++ {
		if err := client.PFAdd(ctx, "myset", fmt.Sprint(i)).Err(); err != nil {
			panic(err)
		}
	}

	card, err := client.PFCount(ctx, "myset").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("set cardinality", card)
}
