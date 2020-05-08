// Run using go run -race 05_atomic.go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	// Prints the machine information
	fmt.Println("#CPU:", runtime.NumCPU())
	fmt.Println("#Goroutines:", runtime.NumGoroutine())

	// It will be a atomic global variable for all the goroutines (Communicate by sharing)
	var counter int64

	// The number of goroutines that will be executed
	const subroutines = 10

	// The `WaitGroup` that will take control over the goroutines
	var wg sync.WaitGroup

	// We add the number of goroutines we want the `WaitGroup` will wait for
	wg.Add(subroutines)

	// This loop will create all the goroutines
	for i := 0; i < subroutines; i++ {
		// The goroutine execution as anonymous function inside the loop
		go func() {
			// It increments by 1 the atomic global variable `counter`, passing its address location
			atomic.AddInt64(&counter, 1)

			// Prints the value of the atomic global variable `counter`, passing its address location
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))

			// It works because there's any lock
			runtime.Gosched()

			// The method `Done` decrements the `WaitGroup` counter, then the anonymous function exits
			wg.Done()
		}()
	}

	// Blocks the main goroutine until the `WaitGroup` counter is 0
	wg.Wait()

	// It prints only 1 goroutine (main) because all the anonymous goroutines exit before the call to `Wait()`
	fmt.Println("#Goroutines:", runtime.NumGoroutine())

	// There's no data race because the atomic global variable avoid it
	// You can now access it without its address location because it is only accessed by the main routine
	fmt.Println("No data race:", counter)
}
