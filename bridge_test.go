package chango

import "fmt"

func ExampleBridge() {
	ch := make(chan (<-chan int))
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			s := make(chan int, i)
			ch <- s
			s <- i
			close(s)
		}
	}()

	done := make(chan interface{})
	defer close(done)

	for v := range Bridge(done, ch) {
		fmt.Printf("%v ", v)
	}
	// Output:
	// 0 1 2 3 4 5 6 7 8 9
}
