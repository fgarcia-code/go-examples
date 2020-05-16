package main

import "fmt"
import "sort"

func main() {
	// This is an unordered slice of int
	int_s := []int{9, 3, 8, 2, 1, 5}

	// This is an unordered slice of string
	string_s := []string{"Francisco", "Benito", "Javier", "Liliana", "Adolfo", "Carlos", "Manuel", "Sarahi"}

	// Printing and sorting the slice of int
	fmt.Println("\nThis is a slice of int unordered:", int_s)
	sort.Ints(int_s)
	fmt.Println("\nThis is a slice of int ordered:", int_s)

	// Printing and sorting the slice of string
	fmt.Println("\nThis is a slice of string unordered:", string_s)
	sort.Strings(string_s)
	fmt.Println("\nThis is a slice of string ordered:", string_s)
}
