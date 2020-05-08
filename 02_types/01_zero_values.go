package main

import "fmt"

func main() {

	// Variables declaration
	var w bool
	var x string
	var y int
	var z float32

	// Default value of a boolean variable
	fmt.Printf("\nThe default value of 'w' is \"%t\" and it's type is \"%T\"\n", w, w)
	w = true
	fmt.Printf("The current value of 'w' is \"%t\" and it's type is \"%T\"\n\n", w, w)

	// Default value of a string variable
	fmt.Printf("The default value of 'x' is \"%s\" and it's type is \"%T\"\n", x, x)
	x = "We change de value"
	fmt.Printf("The current value of 'x' is \"%s\" and it's type is \"%T\"\n\n", x, x)

	// Default value of a int variable
	fmt.Printf("The default value of 'y' is \"%d\" and it's type is \"%T\"\n", y, y)
	y = 10
	fmt.Printf("The current value of 'y' is \"%d\" and it's type is \"%T\"\n\n", y, y)

	// Default value of a float variable
	fmt.Printf("The default value of 'z' is \"%f\" and it's type is \"%T\"\n", z, z)
	z = 10
	fmt.Printf("The current value of 'z' is \"%f\" and it's type is \"%T\"\n\n", z, z)
}