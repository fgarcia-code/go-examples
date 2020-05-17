package main

import "fmt"

func main() {
	// It creates a bidirectional channel
	c := make(chan int)

	// It will create a goroutine, passing the channel as an argument
	go sendFunction(c)

	// It will invoke the function in the main goroutine passing the channel as an argument
	receiveFunction(c)

	// It will print the end of the program
	fmt.Println("The program will exit")
}

// It defines a function with one parameter, this parameter receive a unidirectional channel
// This function only allows sender channel as parameter
func sendFunction(c chan<- int) {
	// The unidirectional channel only allows send data through it
	c <- 10
}

// It defines a function with one parameter, this parameter receive a unidirectional channel
// This function only allows receiver channel as parameter
func receiveFunction(c <-chan int) {
	// The unidirectional channel only allows receive data through it
	fmt.Println("The value received is:", <-c)
}