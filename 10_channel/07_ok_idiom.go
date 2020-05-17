package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 10
		// If you don't close your channel you will get a deadlock
		close(c)
	}()

	v, ok := <-c

	// It print true on the `ok` because the value has been received
	fmt.Println(v, ok)

	v, ok = <-c

	// It print false on the `ok` and the zero value on `v` because the channel is closed
	fmt.Println(v, ok)
}