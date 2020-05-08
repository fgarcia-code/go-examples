package main

import "fmt"

func main() {
	f := calledFunction() // Calling a function that returns a function
	n := f() // Calling the function returned
	fmt.Println(n) // Printing the returned value by the returned function
}

func calledFunction() func() int { // First function with a function as a return
	return func() int { // Returning this function
		n, _ := fmt.Println("Hello I'm a returned function")
		return n // Return this value when the returned function is called
	}
}