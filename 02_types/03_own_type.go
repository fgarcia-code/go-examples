package main

import "fmt"

type custom_type int

func main() {
	a := 42
	var b custom_type = 43
	fmt.Printf("\nValue of 'a' is %v, and it's type is %T\n",a, a)
	fmt.Printf("\nValue of 'b' is %v, and it's type is %T\n\n",b, b)

	// You can not do b = a because these are of different types
	fmt.Printf("You can't do b = a because they are of different types\n\n")

	// You can convert the type 'int' to a 'custom_type'
	b = custom_type(a)
}
