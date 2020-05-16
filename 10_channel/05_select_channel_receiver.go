package main

import "fmt"

func main() {
	// It creates 3 `Channels` 1 for even, 1 for odd, 1 for quit values
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// This goroutine will send values through the `Channel`
	go send(even, odd, quit)

	// this function will receive values through the `Channel`
	selectFunction(even, odd, quit)

	// It finishes gracefully because there are the same number of senders as the number of receivers and nothing stay locked
}

// This function only receives data from the `Channel`
func selectFunction(e, o, q <-chan int) {
	// This is a infinite loop, it finish with a return
	for {
		// The select keyword works with `Channels` it receives data from different `Channel` in each case
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

// This function only send data through the `Channel`
func send(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			// It sends data to the even `Channel`
			e <- i
		} else {
			// It sends data to the odd `Channel`
			o <- i
		}
	}
	// It sends data to the quit `Channel` and it finishes the communication
	q <- -1
}