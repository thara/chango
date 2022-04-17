package chango

import (
	"context"
	"sync"
)

func FanIn[T any](ctx context.Context, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	ch := make(chan T)

	multiplex := func(src <-chan T) {
		defer wg.Done()
		for v := range src {
			select {
			case <-ctx.Done():
			case ch <- v:
			}
		}
	}

	wg.Add(len(channels))
	for _, src := range channels {
		go multiplex(src)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}
