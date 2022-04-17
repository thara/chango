package chango

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func ExampleTake() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	rand.Seed(1)
	rand := func() int { return rand.Int() }

	for n := range Take(ctx, RepeatFn(ctx, rand), 10) {
		fmt.Println(n)
	}
	// Output:
	// 5577006791947779410
	// 8674665223082153551
	// 6129484611666145821
	// 4037200794235010051
	// 3916589616287113937
	// 6334824724549167320
	// 605394647632969758
	// 1443635317331776148
	// 894385949183117216
	// 2775422040480279449
}
