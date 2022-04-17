package chango

func Map[T any, D any](done <-chan D, src <-chan T, f func(T) T) <-chan T {
	dst := make(chan T)
	go func() {
		defer close(dst)
		for v := range src {
			select {
			case <-done:
				return
			case dst <- f(v):
			}
		}
	}()
	return dst
}
