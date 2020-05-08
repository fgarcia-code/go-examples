package main

import "fmt"

func main() {
	for i := 0; i <= 6; i++ {
		fmt.Printf("Factorial of %d is: %d\n",i,factorial(i))
	}

}

//Recursive function (Factorial)
func factorial(n int) int {
	if n == 0 {
		return 1 // The base case
	}
	return n * factorial(n-1) // The function calls itself until a base case
}
