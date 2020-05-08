package main

import "fmt"

var a int
var b int
var s = "Hello Go"

func main() {
	// Empty for
	// for {}
	fmt.Println("\nEmpty for")
	for {
		a++
		fmt.Println(a)
		break
	}

	// Single condition
	// for condition {}
	fmt.Println("\nSingle condition")
	for b <= 2 {
		fmt.Println(b)
		b++
	}

	// For clause
	// for init; condition; post {}
	fmt.Println("\nFor clause")
	for i := 0; i <= 2; i++ {
		fmt.Println(i)
	}


	// Range clause
	// for expression_list := range expression {}
	fmt.Println("\nRange clause")
	for i := range s {
		fmt.Printf("%c\n",s[i])
	}
}