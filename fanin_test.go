package chango

import (
	"context"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func BenchmarkFanIn(b *testing.B) {
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

	for n := 0; n < b.N; n++ {
		for range Take(ctx, FanIn(ctx, fanout...), 10) {
		}
	}
}
