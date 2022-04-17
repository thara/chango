package chango

func OrDone[T any, D any](done <-chan D, src <-chan T) <-chan T {
	dst := make(chan T)
	go func() {
		defer close(dst)
		for {
			select {
			case <-done:
				return
			case v, ok := <-src:
				if !ok {
					return
				}
				select {
				case <-done:
				case dst <- v:
				}
			}
		}
	}()
	return dst
}
