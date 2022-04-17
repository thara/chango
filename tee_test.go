package chango

import (
	"context"
	"fmt"
	"time"
)

func ExampleTee() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	out1, out2 := Tee(ctx, Take(ctx, Repeat(ctx, 1, 2), 4))

	for v := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", v, <-out2)
	}
	// Output:
	// out1: 1, out2: 1
	// out1: 2, out2: 2
	// out1: 1, out2: 1
	// out1: 2, out2: 2
}
