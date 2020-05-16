package main

import "fmt"

func main() {
	// It creates a bidirectional `Channel`
	c := make(chan int)

	// It will create a goroutine, passing the `Channel` as an argument
	go sendFunction(c)

	// It will invoke the function in the `main` goroutine passing the `Channel` as an argument
	receiveFunction(c)

	// It will print the end of the program
	fmt.Println("The program will exit")
}

// It defines a function with one parameter, this parameter receive a unidirectional channel
// This function only allows sender `Channel` as parameter
func sendFunction(c chan<- int) {
	// The unidirectional `Channel` only allows send data through it
	c <- 10
}

// It defines a function with one parameter, this parameter receive a unidirectional channel
// This function only allows receiver `Channel` as parameter
func receiveFunction(c <-chan int) {
	// The unidirectional `Channel` only allows receive data through it
	fmt.Println("The value received is:", <-c)
}