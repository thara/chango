package chango

func Bridge[T any, D any](done <-chan D, ch <-chan <-chan T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for {
			var c <-chan T
			select {
			case <-done:
				return
			case s, ok := <-ch:
				if !ok {
					return
				}
				c = s
			}
			for v := range OrDone(done, c) {
				select {
				case <-done:
				case out <- v:
				}
			}
		}
	}()
	return out
}
