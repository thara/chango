package chango

import (
	"context"
	"math/rand"
	"testing"
	"time"
)

func TestOrDone(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rand := func() int { return rand.Int() }
	src := RepeatFn(ctx.Done(), rand)

	for range OrDone(ctx.Done(), src) {
	}
}
