package main

import "fmt"

func main() {
	message_1 := "Text with double quotes: Hello, go" // Declaration and assign value with double quotes
	message_2 := `Text with back-text: "Hello 
	go in two lines"` // Declaration and assign value with back-text

	// Print strings
	fmt.Println(message_1)
	fmt.Println(message_2)

	byte_slide := []byte(message_1)
	fmt.Printf("Value of message_1 in bytes: `%v`, Type of the message: `%T`\n", byte_slide, byte_slide)
}
