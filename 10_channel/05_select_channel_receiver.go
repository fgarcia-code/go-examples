package main

import "fmt"

func main() {
	// It creates 3 channels 1 for even, 1 for odd, 1 for quit values
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// This goroutine will send values through the channel
	go send(even, odd, quit)

	// This function will receive values through the channel
	selectFunction(even, odd, quit)

	// It finishes gracefully because there are the same number of senders as the number of receivers and nothing stay blocked
}

// This function only receives data from the channel
func selectFunction(e, o, q <-chan int) {
	// This is a infinite loop, it finish with a return
	for {
		// The select keyword works with channels it receives data from different channel in each case
		select {
		case v := <-e:
			fmt.Println("Even value received:", v)
		case v := <-o:
			fmt.Println("Odd value received:", v)
		case v := <-q:
			fmt.Println("Neither even nor odd:", v)
			// It finishes the for loop
			return
		}
	}
}

// This function only send data through the channel
func send(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			// It sends data to the even channel
			e <- i
		} else {
			// It sends data to the odd channel
			o <- i
		}
	}
	// It sends data to the quit channel and it finishes the communication
	q <- -1
}