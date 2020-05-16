package main

import (
	"fmt"
	"sync"
)

// Concurrent fan-in data flow
//                                           |<--- channel1 <--- `x()` <--- for(10-15)
//   `main()` <--- output <--- `merge()` <---|
//                                           |<--- channel2 <--- `y()` <--- for(20-25)
//
func main() {
	// Read from multiple inputs and proceed until all are closed
	// Multiplexing the input channels onto a single channel
	for v := range merge(x(), y()) {
		fmt.Println("I'm reading from output the value:", v)
	}
}

// Input `x()` iterate from 10 to 15
func x() <-chan int {
	channel1 := make(chan int)
	go func() {
		for i := 10; i <= 15; i++ {
			channel1 <- i
		}
		// It close the input `x()`
		close(channel1)
	}()
	return channel1
}

// Input `y()` iterate from 20 to 25
func y() <-chan int {
	channel2 := make(chan int)
	go func() {
		for i := 20; i <= 25; i++ {
			channel2 <- i
		}
		// It close the input `y()`
		close(channel2)
	}()
	return channel2
}

// It merge the inputs `x()` and `y()`
func merge(x, y <-chan int) <-chan int {
	// The multiplexed channel
	output := make(chan int)

	// The WaitGroup that waits for all goroutines to be done
	var wg sync.WaitGroup
	wg.Add(2)

	// goroutine that reads the input `x`, and writes into the single channel
	go func() {
		for input := range x {
			output <- input
		}
		wg.Done()
	}()

	// goroutine that reads the input `y`, and writes into the single channel
	go func() {
		for input := range y {
			output <- input
		}
		wg.Done()
	}()

	// goroutine that waits until all values from  `x` and `y` have been read, then close the output channel
	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}
