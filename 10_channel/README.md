# Channels
**Share memory by communicating**
## Concepts
* A channel provides a **mechanism** for **concurrently** executing **functions** to **communicate** by **sending** and **receiving** **values** of a **specified element type**
* Channels act as **first-in-first-out** queues
### Channel declaration
* A new, initialized channel value can be made using the **built-in** function **make**, which takes the **channel type** and an **optional capacity** as arguments.
  ```go
  // Channel using short variable declaration
  channel := make(chan Type)
  ```
### Send statement
* A send statement **sends** a value **on a channel**
* **Communication blocks** until **the send can proceed** 
* **A send** on a **nil** channel **blocks forever**
  ```go
  // Send statement
  channel <- value
  ```
### Receive operator
* A receive operator **receives** a value **from a channel** 
* The **expression blocks** until a **value** is **available**
* **Receiving** from a **nil** channel **blocks forever**
  ```go
  // Receive operator using short variable declaration
  receiver := <-channel

  // Receive operator using assigment
  receiver <-channel

  // Receive operator ignoring the value
  <-channel
  ```
### Close
* The built-in function **close(c)** records that **no more values** will be **sent** on the **channel** 
* It is an **error** if c is a **receive-only** channel
* **Sending to** or **closing** a **closed channel** causes a **run-time panic**
* **Closing** the **nil** channel also causes a **run-time panic**
* After calling **close**, and after **any previously sent values have been received**, **receive operations** will **return** the **zero value** for the channel's type **without blocking**
  ```go
  // Closing a channel
  close(channel)
  ```
## Buffer Channels
* The **capacity**, in number of elements, sets the **size** of the **buffer** in the **channel**
* If the **capacity** is **zero** or **absent**, the **channel** is **unbuffered** and **communication succeeds only when both a sender and receiver are ready**
* Otherwise, the **channel** is **buffered** and **communication succeeds without blocking** if **the buffer is not full (sends)** or **not empty (receives)**
* A **nil** channel is **never ready** for communication
  ```go
  // Buffer channel channel declaration
  channel := make(chan Type, capacity)
  ```
## Directional Channels
* The optional **<- operator** specifies the **channel direction**, **send** or **receive**. If **no direction** is given, the channel is **bidirectional**. 
  ```go
  // Bidirectional channel using short variable declaration
  channel := make(chan Type)
  
  // Send only channel using short variable declaration
  channel := make(chan <- Type)
  
  // Receive only channel using short variable declaration
  channel := make(<-chan Type)
  
  // Function declaration with receiver channel as return
  func functionName(p1, p2, ..., pn Type1) <-chan Type2 { }
  ```
## Range
* The **iteration values** produced are **the successive values sent** on the channel **until the channel is closed**
* If the **channel** is **nil**, the range expression **blocks forever**
  ```go
  // For channel statement
  for v := range channel {}
  ```
## Select
* A "select" statement **chooses** which of a **set** of **possible send or receive operations** will **proceed**
* **A case** may **assign** the result of **an expression** to **one** or **two variables**, which may be **declared** using a **short variable declaration**
* There can be **at most one default case** and it may appear anywhere in the list of cases.
* Since communication on **nil** channels can **never proceed**, a **select** with **only nil** channels and **no default** case **blocks forever**
  ```go
  // Select statement
  select {
  case <-channel_1:
  ...
  case v := <-channel_2:
  ...
  case channel_3 <- value_1
  ...
  case channel_n <- value_n
  default
  ...
  }
  ```
**Execution:**
1. **All cases are evaluated exactly once**, in source order, **the result is a set of channels** to receive from or send to, and the corresponding values to send. Expressions with a short variable declaration or assignment are not yet evaluated.
2. If one or more of the communications can proceed, **a single one** that **can proceed** is **chosen via** a **uniform pseudo-random selection**. Otherwise, if there is a default case, that case is chosen. If there is **no default case**, **the "select" statement blocks** until **at least one** of the communications can **proceed**.
3. The respective communication operation is executed.
4. If the **selected case** is a **short variable declaration** or **an assignment**, **the expressions** are **evaluated** and **the received value** (or values) are **assigned**.
5. **The statement list** of the selected case **is executed**.
## ok Idiom
* The value of **ok** is **true** if the **value received** was delivered by a **successful send** operation to the channel, or **false** if it is a **zero value** generated because the channel is **closed** and **empty**.
  ```go
  // ok channel idiom
  v, ok := <-channel
  ```
## Fan-In
**A function** can read from **multiple inputs** and proceed **until** all are **closed** by **multiplexing** the **input channels** onto a **single channel** that's **closed** when **all the inputs are closed**. This is called fan-in. [See the example](08_fan_in.go).
```
  input_1 -> Channel_1 ---\
  input_2 -> Channel_2 ----\
  ...                       ---> Channel_out --->
  ...                      /
  input_n -> Channel_n ---/
```
## Fan-Out
**Multiple functions** can read from the **same channel** **until** that channel **is closed** by **demultiplexing** the **input channel** onto a **set of channels** these're **closed** when **the input channel is closed**; this is called fan-out. This provides a way to distribute work amongst a group of workers to parallelize CPU use and I/O. [See the example](09_fan_out.go)
```
                                     /---> Channel_output_1 --->
                                    /----> Channel_output_2 --->
  single_input -> Channel_input --->       ...
                                    \      ...
                                     \---> Channel_output_n --->
```
## Context
* Package context defines the **Context type**, which carries **deadlines**, **cancellation signals**, and other **request-scoped values** across **API boundaries** and **between processes**
  ```go
  // The Context interface in the package context
  type Context interface {
      Deadline() (deadline time.Time, ok bool)
      Done() <-chan struct{}
      Err() error
      Value(key interface{}) interface{}
  }
  ```
* One use case is for server applications when **incoming requests** to the server should **create a Context**, and **outgoing** calls to **another servers** should **accept a Context**
  ```go
  // Simple incoming request
  func incomingRequest(w http.ResponseWriter, req *http.Request){ 
      ctx := context.Context() // It doesn't work is just to demonstrate the incoming request
  }
 
  // Simple outgoing request function
  func outgoingCall(ctx context.Context) { }
  ```
* **The chain of function** calls between them must **propagate the Context**, **optionally replacing** it with a **derived Context** created using WithCancel, WithDeadline, WithTimeout, or WithValue
  ```go
  func function1(){
      ctx := context.Context()
      function2(ctx)   // Chain of function propagates the context
  }
  
  func function2(ctx context.Context) {
      child_ctx := context.WithCancel(ctx) // Replacing ctx with a derived context
      function3(child_ctx)   // Chain of function propagates the context
  }
  
  func fuction3(ctx context.Context) {
      child_ctx := context.WithTimeout(ctx, timeout) // Replacing ctx with a derived context
  }
  ```
* When a **Context is canceled**, **all Contexts derived** from it are **also canceled**
* The WithCancel, WithDeadline, and WithTimeout **functions** take a **Context (the parent)** and **return** a **derived Context (the child)** and a **CancelFunc** 
  ```go
  parent_ctx := context.Context() // Parent context
  child_1_ctx, cancel_1 := context.WithCancel(parent_ctx) // Derived Context
  child_2_ctx, cancel_2 := context.WithDeadline(chile_1_ctx, deadline) // Derived Context
  cancel_1() // It cancels child_1_ctx and child_2_ctx
  ``` 
### Parent Context
The function `context.Context()` in the previous examples doesn't exist because the `context.Context` is an interface, to create a root parent Context you can create it with the functions `context.Background()` and `context.TODO()` as explained below.
#### Background Context
* Background **returns** a non-nil, **empty Context**. It is **never** **canceled**, has **no values**, and has **no deadline** 
* It is typically used by the main function, initialization, and tests, and as the **top-level Context** for incoming requests
  ```go
  ctx := context.Background() // It returns an value of type `*context.emptyCtx`
  
  fmt.Printf("%T\n", ctx) // The type of the context

  output:
  *context.emptyCtx 
  ```
#### TODO Context
* TODO **returns** a non-nil, **empty Context**. **Code should use context**
* TODO **when it's unclear which Context to use** or it is not yet available (because the surrounding function has not yet been extended to accept a Context parameter)
* **TODO is recognized by static analysis tools** that determine whether Contexts are propagated correctly in a program
* **Note:** The best way to start with contexts is using `context.Background()` instead
```go
ctx := context.TODO() // It returns an value of type `*context.emptyCtx`

fmt.Printf("%T\n", ctx) // The type of the context

output:
*context.emptyCtx 
```
### WithCancel Context
* WithCancel **returns a copy of parent** with a new **Done channel**
* The returned context's **Done channel is closed** when the returned **cancel function is called** or when **the parent context's Done channel is closed**, whichever happens first.
* **Canceling** this context **releases resources** associated with it, so code should **call cancel** as soon as **the operations running** in this Context **complete**. [See the example](10_context_cancel.go)
```go
ctx, cancel := context.WithCancel(context.Background()) // It returns a value of type `*context.cancelCtx`

fmt.Printf("%T\n", ctx) // The type of the context
fmt.Printf("%T\n", ctx.Done()) // `ctx.Done()` returns a value of type `<-chan struct {}`
cancel() // It close the channel returned by `ctx.Done()` releasing resources 

output:
*context.cancelCtx
<-chan struct {} 
```
### WithDeadline Context
* **Deadline is a specific time**, that is a time with date, hour, minute and second.
* WithDeadline **returns a copy of the parent** context with **the deadline** adjusted to be **no later than d**. 
* If the parent's **deadline is already earlier than d**, WithDeadline(parent, d) is **semantically equivalent to parent**.
* The returned context's **Done channel is closed** when **the deadline expires**, when **the returned cancel function is called**, or when **the parent context's Done channel is closed**, whichever happens first.
* **Canceling** this context **releases resources** associated with it, so code should **call cancel** as soon as **the operations running** in this Context **complete**. [See the example](10_context_cancel.go)
```go
d := time.Now().Add(time.Second) // The current time plus 1 second as deadline
ctx, cancel := context.WithDeadline(context.Background(), d) // It returns a value of type `*context.timerCtx`

fmt.Printf("%T\n", ctx) // The type of the context
fmt.Printf("%T\n", ctx.Done()) // `ctx.Done()` returns a value of type `<-chan struct {}`

time.Sleep(2 * time.Second) // It waits 2 seconds
fmt.Println(ctx.Err()) // Deadline is reached the channel returned by `ctx.Done()` is also closed with error

cancel() // Cancel has no effect because the channel `ctx.Done()` is already closed

output:
*context.timerCtx
<-chan struct {}
context deadline exceeded
```
### Withtimeout Context
### WithValue Context


