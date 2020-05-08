# Concurrency
## Concurrency vs Parallelism
* Parallelism is about how many tasks you could run in your CPU. If you have 1 core you can't run 2 or more threads in parallel (at the same time)
* Concurrency is a design pattern, it allows run multiple tasks, if you don't have more than one CPU it does't warranty parallelism 
## Concurrency in Go
* **Do not communicate by sharing memory**: 
  * Concurrent programming in many environments is made **difficult** by the subtleties required to implementing the **correct access** to **shared variables**
* **Share memory by communicating**: 
  * Go encourages a different approach in which **shared values** are **passed (sent)** around on channels and, in fact, **never actively shared** by separate threads of execution
  * **Only one goroutine** has **access** to the **value** at any given time
  * [Data races](https://en.wikipedia.org/wiki/Race_condition#Data_race) cannot occur 
  * Go's approach to concurrency originates in Hoare's [Communicating Sequential Processes (CSP)](https://en.wikipedia.org/wiki/Communicating_sequential_processes)
## Goroutines
* They're called **goroutines** because the existing **terms**—**threads**, **coroutines**, **processes**, and so on—convey **inaccurate connotations**
* A **goroutine** has a **simple model**: **it is a function** executing **concurrently** with **other goroutines** in the **same address space** 
* It is **lightweight**, **costing little more** than the allocation of **stack space**
* The stacks **start small**, so they are cheap, and **grow** by allocating (and **freeing**) heap storage **as required**
* **Goroutines** are **multiplexed** onto multiple **OS threads** so if **one should block**, such as while waiting for I/O, **others continue to run**
* Their design hides many of the complexities of thread creation and management
* **Prefix** a **function** or **method** **call** with the `go` keyword **to run** the call in a **new goroutine**. When the **call completes**, the goroutine **exits, silently**
  ```go
  // Starts a new goroutine, and it is executed concurrently
  go f(x, y, z)

  // It happens in the current goroutine (main) 
  f(x,y,z)
  ```
* **Function literals are closures**: the implementation makes sure the **variables** referred to by the function **survive** as long as they are **active**
## Gosched
Gosched [yields](https://en.wikipedia.org/wiki/Yield_(multithreading)) the processor, allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically
* Invocation
  ```go
  runtime.Gosched()
  ```
## WaitGroup
WaitGroup helps to take control over the concurrent *goroutines*
* `WaitGroup` declaration
  ```go
  var wg sync.WaitGroup
  ```
* Make a function concurrent
  ```go
  go functionX()
  ```
* Add the number of *goroutines* to wait
  ```go
  wg.Add(N)
  ```
* Indicate inside each *goroutine* function that it has being done
  ```go
  wg.Done()
  ``` 
* Indicate in the `main` *goroutine* where to wait for the other *goroutines*
  ```go
  wg.Wait()
  ```
## Mutex
`Mutex` helps to define a **block of code** to be **executed** in **mutual exclusion** by **surrounding** it with a call to **Lock** and **Unlock**, in other words, only one *goroutine* will have access to the mutex's lock, locking the surrounding block of code

**Note**: Mutex helps to avoid data race

* Declaration:
  ```go
  var mutex sync.Mutex
  ```
* Taking the mutex's lock
  ```go
  mutex.Lock()
  ```
* Releasing the mutex's lock
  ```go
  mutex.Unlock()
  ```
## Atomic
Package atomic provides **low-level atomic memory primitives** useful for **implementing synchronization algorithms**

These functions **require great care** to be used correctly. Except for special, low-level applications, **synchronization** is **better done with channels** or the facilities of the **sync package**. **Share memory by communicating**; don't communicate by sharing memory.

**Note**: Package atomic helps to avoid data race

**Example for `int64`**:
* Atomic global variable declaration:
  ```go
  var counter int64
  ```
* Accessing by adding `N` to `counter`, inside concurrent *goroutines*:
  ```go
  atomic.AddInt64(&counter,N)
  ```
* Reading value from `counter`, inside councurrent *goroutines*
  ```go
  fmt.Println("Counter\t", atomic.LoadInt64(&counter))
  ```
* Reading value from `counter`, outside concurrent *goroutines*
  ```go
  fmt.Println(counter)
  ```