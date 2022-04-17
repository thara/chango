package chango

import "fmt"

func ExampleTee() {
	done := make(chan interface{})
	defer close(done)

	out1, out2 := Tee(done, Take(done, Repeat(done, 1, 2), 4))

	for v := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", v, <-out2)
	}
	// Output:
	// out1: 1, out2: 1
	// out1: 2, out2: 2
	// out1: 1, out2: 1
	// out1: 2, out2: 2
}
