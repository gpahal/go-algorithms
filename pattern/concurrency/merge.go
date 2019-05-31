package concurrency

import "sync"

// Merge multiplexes a list of channels onto a single channel that is closed when all input
// channels are closed. The ordering for a single channel is maintained but values coming from
// different channels are not ordered in any way.
//
// # Example
//
//     c1 := make(chan int)
//     c2 := make(chan int)
//
//     go func() {
//         c1 <- 1
//         c2 <- 2
//     }
//
//     for n := range concurrency.Merge(c2, c2) {
//         fmt.Println(n) // prints 1 and 2, or 2 and 1
//     }
func Merge(cs ...<-chan int) <-chan int {
	out := make(chan int)

	// A WaitGroup is used here to wait for len(cs) goroutines to terminate. These goroutines are
	// started in the for the for statement below for each input channel.
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		// Start a goroutine for each input channel. It copies values from its input channel to the
		// out channel until the input channel is closed. Then the goroutine calls wg.Done to
		// indicate it is about to terminate.
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}

	// Start a goroutine to close the out channel once all the output goroutines are done.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
