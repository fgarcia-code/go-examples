package main

import "fmt"

func main() {
	// It will create a new channel
	c := make(chan int)

	go func() {
		// It will iterate sending data over the channel
		for i := 0; i < 10; i++ {
			// It sends the value i over the channel
			c <- i
		}
		// You need to close the channel, if not the for in the main goroutine will be blocked, and you'll get a deadlock
		close(c)
	}()

	// We iterate over all values received in the channel until it gets closed, if you don't close it, the main goroutine will be blocked
	for v := range c {
		fmt.Println("The new value received is:", v)
	}
}