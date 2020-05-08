package main

import "fmt"

// Global declaration
var y = 43

// The identifiers of the varibles are x, y and z
func main() {
	// Short operator declaration
	x := 27
	x = 24
	fmt.Println("Just local variable:", x)

	// Let's print global variable
	println("Global variable:", y)

	// Local varible with var with type
	var z int = 10
	println("Typed variable:",z)
}
