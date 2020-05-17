package main

import "fmt"

func main() {
	// It create 2 channel 1 for values an 1 for exit
	add := make(chan int)
	quit := make(chan int)

	// It creates a sender goroutine
	go sender(add, quit)

	// It will invoke a function receiver
	receiver(add, quit)
}

// This function has 1 parameter `a` that sends data and 1 paremeter `q` that receives data
func sender(a chan<- int, q <-chan int) {
	x := 0
	// This is a infinite loop
	for {
		select {
		// It sends data
		case a <- x:
			x++
	    // It receives data, and it will break the infinite loop
		case <-q:
			return
		}
	}
}

// This function has 1 parameter `a` that receives data and 1 parameter `q` that sends data
func receiver(a <-chan int, q chan<- int) {
	for i := 0; i < 10; i++ {
		// This will print the data received
		fmt.Println("The next value is:", <-a)
	}
	// This will send the exit signal to the infinite loop
	q <- 0
}