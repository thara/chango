package chango

import "context"

func Bridge[T any](ctx context.Context, ch <-chan <-chan T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for {
			var c <-chan T
			select {
			case <-ctx.Done():
				return
			case s, ok := <-ch:
				if !ok {
					return
				}
				c = s
			}
			for v := range OrDone(ctx, c) {
				select {
				case <-ctx.Done():
				case out <- v:
				}
			}
		}
	}()
	return out
}
