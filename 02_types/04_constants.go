package main

import (
	"fmt"
)

func main() {

	// Constants declaration
	const (
		a = 41
		b = 42.72
		c = "This is a constant"
	)

	fmt.Printf("\nThe constant: %v, hast the type: %T\n", a, a)
	fmt.Printf("The constant: %v, hast the type: %T\n", b, b)
	fmt.Printf("The constant: %v, hast the type: %T\n\n", c, c)

	// Iota constants
	const (
		d = iota
		e
		f
	)

	// Iota constants
	const (
		g = iota
		h
		i
	)

	fmt.Println("Constants with iota:")
	fmt.Println("d is:",d)
	fmt.Println("e is:",e)
	fmt.Println("f is:",f)
	fmt.Println("g is:",g)
	fmt.Println("h is:",h)
	fmt.Println("i is:",i)

	// iota with byte shifting
	const (
		_ = iota
		j = 1 << (iota * 1)
		k = 1 << (iota * 1)
		l = 1 << (iota * 1)
	)

	fmt.Println("\nConstant byte shifting with iota: ")
	fmt.Printf("%d\t\t%b\n",j,j)
	fmt.Printf("%d\t\t%b\n",k,k)
	fmt.Printf("%d\t\t%b\n\n",l,l)

}
