package chango

import (
	"context"
	"fmt"
	"time"
)

func ExampleMap() {
	mul := func(v int) int {
		return v * 2
	}
	add := func(v int) int {
		return v + 1
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	src := Generator(ctx, 1, 2, 3, 4)
	pipeline := Map(ctx, Map(ctx, Map(ctx, src, mul), add), mul)

	for v := range pipeline {
		fmt.Println(v)
	}
	// Output:
	// 6
	// 10
	// 14
	// 18
}
