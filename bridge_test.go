package chango

import (
	"context"
	"fmt"
	"time"
)

func ExampleBridge() {
	ch := make(chan (<-chan int))
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			s := make(chan int, i)
			ch <- s
			s <- i
			close(s)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for v := range Bridge(ctx, ch) {
		fmt.Printf("%v ", v)
	}
	// Output:
	// 0 1 2 3 4 5 6 7 8 9
}
