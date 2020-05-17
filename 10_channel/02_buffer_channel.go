package main

import "fmt"

func main() {
	// It will create a buffer channel, it allows 2 elements in the buffer
	// The buffer channel will not block the goroutine, as long as, it has only 2 elements (in this case)
	c := make(chan int, 2)

	// It will send the first value into the buffer channel
	c <- 10

	// It will send the second value into the buffer channel
	c <- 11

	// The buffer channel is full, if you send a third value it will block the main goroutine
	// c <- 11

	// It receives values from the buffer channel
	fmt.Println("The first value in the channel is:", <-c)
	fmt.Println("The second value in the channel is:", <-c)

	// Once you have receive values from the buffer channel you can send more values, with out block the main goroutine
	c <- 12

	// It continues receiving values from the buffer channel
	fmt.Println("The third value in the channel is:", <-c)
}