# chango
go generic channel utilities; inspired by Concurrency in Go

## Usage

**generators**
```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

rand.Seed(1)
rand := func() int { return rand.Int() }

for n := range Take(ctx, RepeatFn(ctx, rand), 10) {
	fmt.Println(n)
}
```

**Map**
```go
mul := func(v int) int { return v * 2 }
add := func(v int) int { return v + 1 }

ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

src := Generator(ctx, 1, 2, 3, 4)
pipeline := Map(ctx, Map(ctx, Map(ctx, src, mul), add), mul)
for v := range pipeline {
	fmt.Println(v)
}
// Output:
// 6
// 10
// 14
// 18
```

**OrDone**
```go
ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
defer cancel()

rand := func() int { return rand.Int() }
src := RepeatFn(ctx, rand)

for v := range OrDone(ctx, src) {
  fmt.Println(v)
}
```

**FanIn**
```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

rand := func() int { return rand.Intn(5000000) }
randCh := RepeatFn(ctx, rand)

N := runtime.NumCPU()
fanout := make([]<-chan int, N)
for i := 0; i < N; i++ {
	fanout[i] = primeFinder(ctx, randCh)
}

var result []int
for p := range Take(ctx, FanIn(ctx, fanout...), 10) {
	fmt.Println(p)
}
```

**Tee**
```go
ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
defer cancel()

out1, out2 := Tee(ctx, Take(ctx, Repeat(ctx, 1, 2), 4))

for v := range out1 {
	fmt.Printf("out1: %v, out2: %v\n", v, <-out2)
}
// Output:
// out1: 1, out2: 1
// out1: 2, out2: 2
// out1: 1, out2: 1
// out1: 2, out2: 2
```

**Bridge**
```go
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

ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

for v := range Bridge(ctx, ch) {
	fmt.Printf("%v ", v)
}
// Output:
// 0 1 2 3 4 5 6 7 8 9
```


## Installation

```
go get github.com/thara/chango
```

## License

MIT

## Author

Tomochika Hara (a.k.a [thara](https://thara.dev))
