package chango

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-Or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(4*time.Second),
		sig(3*time.Second),
	)

	if 1.1 <= time.Since(start).Seconds() {
		t.Errorf("done after %v", time.Since(start))
	}
}
