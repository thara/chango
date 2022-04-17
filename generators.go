package chango

import "context"

func Generator[T any](ctx context.Context, values ...T) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, v := range values {
			select {
			case <-ctx.Done():
				return
			case ch <- v:
			}
		}
	}()
	return ch
}

func Repeat[T any](ctx context.Context, values ...T) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for {
			for _, v := range values {
				select {
				case <-ctx.Done():
					return
				case ch <- v:
				}
			}
		}
	}()
	return ch
}

func RepeatFn[T any](ctx context.Context, fn func() T) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- fn():
			}
		}
	}()
	return ch
}

func Take[T any](ctx context.Context, src <-chan T, n int) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- <-src:
			}
		}
	}()
	return ch
}
