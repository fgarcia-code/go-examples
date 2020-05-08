package main

import "fmt"

func main() {
	m := "Hello"

	func(message string){ // function parameters
		fmt.Println(message, "anonymous function")
	}(m) // function arguments
}
