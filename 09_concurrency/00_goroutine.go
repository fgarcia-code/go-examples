package main

import (
	"fmt"
	"runtime"
)

func main() {
	// The number of goroutines is 1 (main goroutine)
	fmt.Println("The number of goroutines are:", runtime.NumGoroutine())

	// We are running a new goroutine prefixing the method Println with de keyword `go`
	// The `Println` output will not be printed, because it's flow runs concurrently
	// The return of every function will be discarded when you make it goroutine
	go fmt.Println("This is another goroutine")

	//  The number of goroutines are 2
	fmt.Println("The number of goroutines are:", runtime.NumGoroutine())
}
