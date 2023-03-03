package main

import (
	"context"
	"fmt"
	"math/rand"

	c "github.com/lzzzzl/gogo-learn/redis"
)

// 需要 RedisBloom 模块支持
//
// git clone --recursive https://github.com/RedisBloom/RedisBloom.git
// cd redisbloom
// make

// 加载实时模块
// /path/to/redis-server --loadmodule ./redisbloom.so

var (
	ctx    = context.Background()
	client = c.Client1()
)

// 布隆过滤器
func bloomFilter(ctx context.Context) {
	inserted, err := client.Do(ctx, "BF.ADD", "bf_key", "item0").Bool()
	if err != nil {
		panic(err)
	}
	if inserted {
		fmt.Println("item0 was interted")
	}

	for _, item := range []string{"item0", "item1"} {
		exists, err := client.Do(ctx, "BF.EXISTS", "bf_key", item).Bool()
		if err != nil {
			panic(err)
		}
		if exists {
			fmt.Printf("%s does exist\n", item)
		} else {
			fmt.Printf("%s does not exist\n", item)
		}
	}
}

// Cuckoo过滤器
func cuckooFilter(ctx context.Context) {
	inserted, err := client.Do(ctx, "CF.ADDNX", "cf_key", "item0").Bool()
	if err != nil {
		panic(err)
	}
	if inserted {
		fmt.Println("item0 was inserted")
	} else {
		fmt.Println("item0 already exists")
	}

	for _, item := range []string{"item0", "item1"} {
		exists, err := client.Do(ctx, "CF.EXISTS", "cf_key", item).Bool()
		if err != nil {
			panic(err)
		}
		if exists {
			fmt.Printf("%s does exist\n", item)
		} else {
			fmt.Printf("%s does not exist\n", item)
		}
	}

	deleted, err := client.Do(ctx, "CF.DEL", "cf_key", "item0").Bool()
	if err != nil {
		panic(err)
	}
	if deleted {
		fmt.Println("item0 was deleted")
	}
}

// 计数最小图
func countMinSketch(ctx context.Context) {
	if err := client.Do(ctx, "CMS.INITBYPROB", "count_min", 0.001, 0.01).Err(); err != nil {
		panic(err)
	}

	items := []string{"item1", "item2", "item3", "item4", "item5"}
	counts := make(map[string]int, len(items))

	for i := 0; i < 10000; i++ {
		n := rand.Intn(len(items))
		item := items[n]

		if err := client.Do(ctx, "CMS.INCRBY", "count_min", item, 1).Err(); err != nil {
			panic(err)
		}
		counts[item]++
	}

	for item, count := range counts {
		ns, err := client.Do(ctx, "CMS.QUERY", "count_min", item).Int64Slice()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s: count-min=%d actual=%d\n", item, ns[0], count)
	}
}

// topK
func topK(ctx context.Context) {
	if err := client.Do(ctx, "TOPK.RESERVE", "top_items", 3).Err(); err != nil {
		panic(err)
	}

	counts := map[string]int{
		"item1": 1000,
		"item2": 2000,
		"item3": 3000,
		"item4": 4000,
		"item5": 5000,
		"item6": 6000,
	}

	for item, count := range counts {
		for i := 0; i < count; i++ {
			if err := client.Do(ctx, "TOPK.INCRBY", "top_items", item, 1).Err(); err != nil {
				panic(err)
			}
		}
	}

	items, err := client.Do(ctx, "TOPK.LIST", "top_items").StringSlice()
	if err != nil {
		panic(err)
	}

	for _, item := range items {
		ns, err := client.Do(ctx, "TOPK.COUNT", "top_items", item).Int64Slice()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s: top-k=%d actual=%d\n", item, ns[0], counts[item])
	}
}
