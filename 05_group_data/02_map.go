package main

import "fmt"

func main() {
	// Declaring a map
	m := map[string]int{
		"Mexico": 1968,
		"USA":    1996,
		"Canada": 1976,
	}

	// Printing the map
	fmt.Println("\nThis is a map:")
	fmt.Println(m)

	// Accessing to an element's map by key
	fmt.Println("\nAccessing to the element 'Mexico':")
	fmt.Println("m[\"Mexico\"]",m["Mexico"])

	// Checking if a key exist in the map
	fmt.Println("\nChecking if key 'China' exist in the map:")
	if _,ok := m["China"]; !ok {
		fmt.Println("The key doesn't exists")
	}

	// Add element
	fmt.Println("\nAdding the element 'Russia':")
	m["Russia"] = 2018
	fmt.Println(m)

	// Iterate though the map
	fmt.Println("\nIterating through the map:")
	for k, v := range m {
		fmt.Println(k,v)
	}

	// Delete an element
	fmt.Println("\nDeleting the element 'USA':")
	delete(m,"USA")
	fmt.Println(m)
	fmt.Println()
}
