package chango

func Generator[T any, D any](done <-chan D, values ...T) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, v := range values {
			select {
			case <-done:
				return
			case ch <- v:
			}
		}
	}()
	return ch
}

func Repeat[T any, D any](done <-chan D, values ...T) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case ch <- v:
				}
			}
		}
	}()
	return ch
}

func RepeatFn[T any, D any](done <-chan D, fn func() T) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			case ch <- fn():
			}
		}
	}()
	return ch
}

func Take[T any, D any](done <-chan D, src <-chan T, n int) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case ch <- <-src:
			}
		}
	}()
	return ch
}
