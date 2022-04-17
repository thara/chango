package chango

func Tee[T any, D any](done <-chan D, src <-chan T) (_, _ <-chan T) {
	out1 := make(chan T)
	out2 := make(chan T)
	go func() {
		defer close(out1)
		defer close(out2)
		for v := range OrDone(done, src) {
			var o1, o2 = out1, out2
			for i := 0; i < 2; i++ {
				select {
				case o1 <- v:
					o1 = nil
				case o2 <- v:
					o2 = nil
				}
			}
		}
	}()
	return out1, out2
}
