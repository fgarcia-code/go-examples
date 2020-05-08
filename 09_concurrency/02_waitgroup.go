package main

import (
	"fmt"
	"runtime"
	"sync"
)

// It creates a `WaitGroup`
var wg sync.WaitGroup

func main() {
	// Print some system information
	fmt.Println("System Information:")
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPU\t\t", runtime.NumCPU(), "\n")

	// Print the current number of goroutines: 1 (main)
	fmt.Println("GOROUTINES\t", runtime.NumGoroutine())

	// `WaitGroup` method Add(delta int)` where `delta` is the number of additional goroutines to wait
	wg.Add(1)

	// Creates a new goroutine
	// It will be executed concurrently
	// If #CPU > 1 it will be executed in parallel (at the same time as the main)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Second:", i)
		}

		// Method `Done()` decrements the `WaitGroup` counter by one
		wg.Done()
	}()

	// This function will be executed in the `main` goroutine
	func() {
		for i := 0; i < 10; i++ {
			fmt.Println("First:", i)
		}
	}()

	// Print the current number of goroutines: 2 (main and goroutine)
	fmt.Println("\nGOROUTINES\t", runtime.NumGoroutine())

	// It waits for the number of `delta` goroutines
	wg.Wait()
}
