package main

import (
	"fmt"
)

func main() {
	a := 10
	// Expression switch comparison
	fmt.Println("\nExpression switch comparison:")
	switch {
		case a < 10:
			fmt.Println("'a' less than 10")
			fallthrough
		case a > 10:
			fmt.Println("'a' greater than 10")
			fallthrough
		case a <= 10:
			fmt.Println("'a' less or equal 10")
			fallthrough
		case a >= 10:
			fmt.Println("'a' greater or equal 10")
			fallthrough
		default:
			fmt.Println("'a' equals 10")
	}

	// Expression switch tag
	fmt.Println("\nExpression switch tag:")
	fmt.Printf("The string: ")
	switch a {
	case 2, 4, 6, 8, 10, 12:
		fmt.Println("'a' is in: 2, 4, 6, 8, 10, 12")
	case 1, 3, 5, 7, 9, 11:
		fmt.Println("'a' is in: 1, 3, 5, 7, 9, 11")
	}

	// Expression switch function
	fmt.Println("\nExpression switch function:")
	fmt.Printf("The string: ")
	switch x, _ := fmt.Println("Hello"); {
	case x < 10:
		fmt.Println("Length less than 10\n")
	case x > 10:
		fmt.Println("Length less than 10\n")
	default:
		fmt.Println("Length equals 10\n")
	}

}
