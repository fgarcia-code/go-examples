package main

import "fmt"

func main() {
	// Print pair numbers with control flow statements
	fmt.Println("\nPrint pair numbers with control flow statements:")
	for i := 0; i <= 10; i++ {
		if (i % 2) != 0 {
			continue
		}
		fmt.Println("The value of 'i' is:", i)
	}
}
