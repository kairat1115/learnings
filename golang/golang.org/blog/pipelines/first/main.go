package main

// Link: https://blog.golang.org/pipelines

import (
	"fmt"
	"sync"
)

// Stage 1.
func gen(nums ...int) <-chan int {
	// Bad code
	// out := make(chan int)

	// we can simplify this function, because length of numbers slice is known
	// this will avoid creating a new goroutine.
	out := make(chan int, len(nums))

	// send all numbers in goroutine
	// and close channel `out` when done sending
	// go func() {
	// 	for _, n := range nums {
	// 		out <- n
	// 	}
	// 	close(out)
	// }()

	for _, n := range nums {
		out <- n
	}
	close(out)

	return out
}

// Stage 2.
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	// get all numbers from channel in
	// square number and send to channel `out`
	// `for` will not be infinite, because
	// channel from previous pipeline is closed.
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// The merge function converts a list of channels to a single channel
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	// out := make(chan int)
	out := make(chan int, 1) // enough space for the unread inputs

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, ch := range cs {
		go output(ch)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	// setting up pipeline
	// similar in shel will be `echo 2 3 | gen | sq`
	ch := gen(2, 3)
	out := sq(ch)

	// print output
	for n := range out {
		fmt.Println(n)
		// 2 * 2 = 4
		// 3 * 3 = 9
	}

	// we can write compactly by passing channel to function directly
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n)
		// 4 * 4 = 16
		// 9 * 9 = 81
	}

	// Multiple functions can read from the same channel until that channel is closed;
	// this is called fan-out.
	// This provides a way to distribute work amongst a group of workers
	// to parallelize CPU use and I/O.

	// setup
	ch = gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	ch1 := sq(ch)
	ch2 := sq(ch)

	for n := range merge(ch1, ch2) {
		fmt.Println(n)
	}
}
