package main

import (
	"fmt"
	"sync"
	"time"
)

// Concurrent fan-out data flow
//                                |---> `a()` ---> output1
//   `generator()` ---> input --->|
//                                |---> `b()` ---> output2
//
func main() {
	// Multiple functions can read from the same channel until that channel is closed
	input := generator()

	// The WaitGroup that waits for all goroutines to be done
	var wg sync.WaitGroup
	wg.Add(2)

	// Reading `input` channel from two different `output` channels
	output1 := a(input)
	output2 := b(input)

	// Reading a sample of the `input` channel from the `output1` channel
	go func() {
		for v := range output1 {
			fmt.Println("I'm reading from output1 the value is:\t", v)
		}
		wg.Done()
	}()

	// Reading a sample of the `input` channel from the `output2` channel
	go func() {
		for v := range output2 {
			fmt.Println("I'm reading from output2 the values is:\t", v)
		}
		wg.Done()
	}()

	// It will wait until all values from `generator()` have been read
	wg.Wait()
}

// It will generate some values in a single `input` channel
func generator() <-chan int {
	input := make(chan int)
	// It will generate numbers from 0 to 20, then it will send each value through the `input` channel
	go func() {
		for i := 0; i <= 20; i++ {
			input <- i
		}

		// It will close the channel when all values have been generated
		close(input)
	}()

	return input
}

// It will read some values from the `input` channel
func a(input <-chan int) <-chan int {
	output1 := make(chan int)
	go func() {
		for v := range input {
			output1 <- v
			time.Sleep(time.Second)
		}
		close(output1)
	}()

	return output1
}

// It will read some values from the `input` channel
func b(input <-chan int) <-chan int {
	output2 := make(chan int)
	go func() {
		for v := range input {
			output2 <- v
			time.Sleep(time.Second)
		}
		close(output2)
	}()

	return output2
}
