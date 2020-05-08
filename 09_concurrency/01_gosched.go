package main

import (
	"fmt"
	"runtime"
)

func main() {
	// The number of goroutines is 1 (main goroutine)
	fmt.Println("The number of goroutines are:", runtime.NumGoroutine())

	// Running an anonymous function as a goroutine
	go func() {
		fmt.Println("This is another goroutine")

		// If you uncomment the `Gosched` below, the goroutine will be paused
		// and the last `Println` will print 2
		//runtime.Gosched()
	}()

    // `Gosched` yields the processor, allowing other goroutines to run
    runtime.Gosched()

	// Anonymous function run and exits thanks to `Gosched()`
    // because of anonymous function has finished, now we just have only 1 goroutine
	fmt.Println("The number of goroutines are:", runtime.NumGoroutine())
}
