package main

import "fmt"

func main()  {
	defer closeFunction() // It will be executed at the end of main
	openFunction()
} // closeFunction will be executed

func openFunction()  {
	fmt.Println("This is the open function")
}

func closeFunction()  {
	fmt.Println("This is the close function")
}