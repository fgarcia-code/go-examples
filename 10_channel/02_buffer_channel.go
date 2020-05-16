package main

import "fmt"

func main() {
	// It will create a buffer `Channel`, it allows 2 elements in the buffer
	// The buffer `Channel` will not lock the goroutine, as long as, it has only 2 elements (in this case)
	c := make(chan int, 2)

	// It will send the first value into the buffer `Channel`
	c <- 10

	// It will send the second value into the buffer `Channel`
	c <- 11

	// The buffer `Channel` is full, if you send a third value it will lock the `main` goroutine
	// c <- 11

	// It receives the values from the buffer `Channel`
	fmt.Println("The first value in the channel is:", <-c)
	fmt.Println("The second value in the channel is:", <-c)

	// Once you have receive the values from the buffer `Channel` you can send more values, with out lock the main routine
	c <- 12

	// It continues receiving values from the buffer `Channel`
	fmt.Println("The third value in the channel is:", <-c)
}