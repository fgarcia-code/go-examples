package main

import "fmt"

func main() {

	// Single condition
	fmt.Println("\nSingle condition")
	if true {
		fmt.Println("Hello Go")
	}

	// If with init statement
	fmt.Println("\nIf with init statement")
	if x := 10; x == 10 {
		fmt.Printf("'x' is %d\n\n",x)
	}

	// If else if statement
	fmt.Println("\nIf else if statement")
	if x := 11; x == 10 {
		fmt.Printf("'x' is equal to 10")
	} else if x <= 10 {
		fmt.Println("'x' is less than 10")
	} else {
		fmt.Println("'x' is greater than 10")
	}
}
