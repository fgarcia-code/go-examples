package main

import "fmt"

func main() {
	x := 255
	// Printf usage
	fmt.Println("\nThe 'x' variable using 'Printf' function is:")
	fmt.Printf("'x' in binary: %b, in hexadecimal: %#x, in decimal: %d\n",x,x,x)

	// Sprintf  usage
	s := fmt.Sprintf("'x' in binary: %b, in hexadecimal: %#x, in decimal: %d\n",x,x,x)
	fmt.Println("\nThe 's' variable using 'Sprintf' function is:")
	fmt.Println(s)
}
