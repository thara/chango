package chango

import (
	"fmt"
)

func ExampleMap() {
	done := make(chan interface{})
	defer close(done)

	mul := func(v int) int {
		return v * 2
	}
	add := func(v int) int {
		return v + 1
	}

	src := Generator(done, 1, 2, 3, 4)
	pipeline := Map(done, Map(done, Map(done, src, mul), add), mul)

	for v := range pipeline {
		fmt.Println(v)
	}
	// Output:
	// 6
	// 10
	// 14
	// 18
}
