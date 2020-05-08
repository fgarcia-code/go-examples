package main

import "fmt"

func main() {
	// x := type{values}
	x := []int{2,4,6,8,10}
	fmt.Printf("\nThe array:\n%v\n\n",x)

	// Iterate through slice
	fmt.Println("Iterate through the elements of the array:")
	for i, v := range x {
		fmt.Println("The index:", i, ", and the value", v)
	}

	// Slicing the slice
	fmt.Println("\nThese are the elements from 1 to 3:", x[1:3])

	// Append to slice
	y := []int{50,60,70}
	x = append(x, 20, 30, 40)
	x = append(x, y...)
	fmt.Println("\nThis is the new slice after appending:",x)

	// Delete from slice
	x = append(x[:2],x[4:]...)
	fmt.Println("\nThis is the new slice after deletion:",x)

	// Make allocation
	fmt.Println("\nSlice with make allocation:")
	z := make([]int, 4, 8)
	fmt.Printf("%v\t\t%d\t%d", z,len(z),cap(z))
	z = append(z,1,2,3,4,5)
	fmt.Printf("\n%v\t%d\t%d\n\n", z,len(z),cap(z))

	// Multi-dimensional slice
	p := [][]string{{"America","Mexico","CDMX"},{"America","United States","Washington"}}
	fmt.Println("This is  a multi-dimensional slice:")
	fmt.Println(p)
}