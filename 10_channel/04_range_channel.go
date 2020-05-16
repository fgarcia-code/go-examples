package main

import "fmt"

func main() {
	// It will create a new channel
	c := make(chan int)

	go func() {
		// It will iterate sending data over the `Channel`
		for i := 0; i < 10; i++ {
			// It sends the value i over the `Channel`
			c <- i
		}
		// You need to close the `Channel`, if not the for in the `main` goroutine will be locked, and you'll get a deadlock
		close(c)
	}()

	// We iterate over all values received in the `Channel` until it gets closed, if you don't close it, the `main` goroutine will be locked
	for v := range c {
		fmt.Println("The new value received is:", v)
	}
}