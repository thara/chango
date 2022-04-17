package chango

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"time"
)

func ExampleFunIn() {
	primeFinder := func(ctx context.Context, src <-chan int) <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for n := range src {
				n -= 1
				prime := true
				for div := n - 1; 1 < div-1; div-- {
					if n%div == 0 {
						prime = false
						break
					}
				}
				if prime {
					select {
					case <-ctx.Done():
						return
					case ch <- n:
					}
				}
			}
		}()
		return ch
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	rand.Seed(1)
	rand := func() int { return rand.Intn(5000000) }
	randCh := RepeatFn(ctx, rand)

	N := runtime.NumCPU()
	fanout := make([]<-chan int, N)
	for i := 0; i < N; i++ {
		fanout[i] = primeFinder(ctx, randCh)
	}

	var result []int
	for p := range Take(ctx, FanIn(ctx, fanout...), 10) {
		result = append(result, p)
	}

	sort.Sort(sort.IntSlice(result))
	for _, v := range result {
		fmt.Println(v)
	}

	// Output:
	// 270059
	// 931967
	// 2107939
	// 2131463
	// 2393161
	// 2694161
	// 2921531
	// 3108509
	// 3221269
	// 3958589
}
