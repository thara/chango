package chango

func Or[T any](ch1, ch2 <-chan T, others ...<-chan T) <-chan T {
	var or func(...<-chan T) <-chan T
	or = func(channels ...<-chan T) <-chan T {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}
		done := make(chan T)
		go func() {
			defer close(done)
			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], done)...):
				}
			}
		}()
		return done
	}

	channels := []<-chan T{}
	channels = append(channels, ch1)
	channels = append(channels, ch2)
	channels = append(channels, others...)
	return or(channels...)
}
