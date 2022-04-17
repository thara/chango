package chango

import "context"

func OrDone[T any](ctx context.Context, src <-chan T) <-chan T {
	dst := make(chan T)
	go func() {
		defer close(dst)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-src:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
				case dst <- v:
				}
			}
		}
	}()
	return dst
}
