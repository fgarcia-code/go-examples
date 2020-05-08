package main

import "fmt"

func main() {
	// Create a variable 'a'
	a := 10

	// Printing the value of 'a'
	fmt.Println("\nThis is the 'a' value:")
	fmt.Println(a)

	// Printing the type of 'a'
	fmt.Printf("This is the 'a' type: %T\n",a)

	// Printing the memory addres of 'a'
	fmt.Println("\nThis is the 'a' memory address:")
	fmt.Println(&a)

	/// Printing the value of poiter to 'a'
	fmt.Printf("This is the 'a' pointer type: %T\n",&a)

	// Creating a pointer '*b'
	b := &a

	// Printing the value of 'b'
	fmt.Println("\nThis is 'b' value:")
	fmt.Println(b)

	// Printing the value of '*b'
	fmt.Println("\nThis is the '*b' value:")
	fmt.Println(*b)

	// Changing the 'a' value
	*b = 20
	fmt.Println("\nThe new 'a' value is:",a)
}
