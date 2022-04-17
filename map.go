package chango

import "context"

func Map[T any](ctx context.Context, src <-chan T, f func(T) T) <-chan T {
	dst := make(chan T)
	go func() {
		defer close(dst)
		for v := range src {
			select {
			case <-ctx.Done():
				return
			case dst <- f(v):
			}
		}
	}()
	return dst
}
