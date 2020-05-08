// Run using go run -race 03_data_race.go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Prints the machine information
	fmt.Println("#CPU:", runtime.NumCPU())
	fmt.Println("#Goroutines:", runtime.NumGoroutine())

	// It will be a global variable for all the goroutines (Communicate by sharing)
	global_count := 0

	// The number of goroutines that will be executed
	const subroutines = 100

	// The `WaitGroup` that will take control over the goroutines
	var wg sync.WaitGroup

	// We add the number of goroutines we want the `WaitGroup` will wait for
	wg.Add(subroutines)

	// This loop will create all the goroutines
	for i := 0; i < subroutines; i++ {
		// The goroutine execution as anonymous function inside the loop
		go func() {
			// The `local_count` scope only exists in the anonymous function
			// `global_count` is shared across all goroutines
			local_count := global_count // "Communicating by sharing"

			// It yields the processor, the goroutine execution is paused
			runtime.Gosched() // It says to the processor "run something else"

			// The goroutine is resumed automatically here
			local_count++

			// "Communicating by sharing"
			global_count = local_count // Global count is modified by all goroutines, causing a data race

			// The method `Done` decrements the `WaitGroup` counter, then the anonymous function exits
			wg.Done()
		}()
	}

	// Blocks the main goroutine until the `WaitGroup` counter is 0
	wg.Wait()

	// It prints only 1 goroutine (main) because all the anonymous goroutines exit before the call to `Wait()`
	fmt.Println("#Goroutines:", runtime.NumGoroutine())

	// Data race condition, unexpected behavior
	fmt.Println("Data race, unexpected behavior:", global_count)
}
