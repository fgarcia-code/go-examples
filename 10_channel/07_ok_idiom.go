package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 10
		// If you dont close your channel you will get a deadlock
		close(c)
	}()

	v, ok := <-c

	// It will print true on the ok because the value has been received
	fmt.Println(v, ok)

	v, ok = <-c

	// It will print false on the ok because it doesn't receive any value
	fmt.Println(c, ok)
}