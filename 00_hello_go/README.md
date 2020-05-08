# Hello Go
`package main`: Declares the main package

`import "fmt"`: Imports the package `fmt`


``` go
// Declares the entry point of th program.
// The `{` must be in the same line as the main

func main() {
}
```

`fmt.Println("Hello, Go")`: Executes the function Println() of the pakage `fmt`

`fmt`: Package fmt implements formatted I/O with functions analogous to C's printf and scanf

`func Print(a ...interface{}) (n int, err error)`: The definition of the function `Print`
 * `...`: Indicate any parameters
 * `interface{}`: Every thing is of type `interface{}`, It is the empty interface.
 * `n`: It is the number of bytes written.
 * `err`: It is the error's number
 * `_`: Indicates don't do anything with the return