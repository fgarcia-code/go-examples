package main

import "fmt"

func main() {
	// Code block scope
	{
		y := 10
		fmt.Println("y only live here:",y)
	}

	// The scope a() is created
	a := incrementor()
	// The scope b() is created
	b := incrementor()

	fmt.Println("\nThis is the a() scope:")
	fmt.Println(a())

	fmt.Println("This is the b() scope:")
	fmt.Println(b())
	fmt.Println(b())
}

// Function scope
func incrementor() func() int {
	var x int // This variable is created
	return func() int {
		x++ // This variable is alive through the function's scope
		return x
	}
}