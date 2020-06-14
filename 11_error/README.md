# Error Handling
## Error
* The type `error` is defined as:
  ```go
  type error interface {
     Error () string
  }
  ```
* Every type that implements the method error is also of type error
  ```go
  type CustomError string

  func (ce CustomError) Error() string {
    return "This is an error"
  }
  
  main()
  ---
  var ce error
  ce = CustomError("")
  fmt.Printf(ce.Error())
  ```
* Go's multivalue return makes it easy  to return a detailed error description alongside the normal return value
  ```go
  main()
  ---
  ce, error := functionError()
  fmt.Println(ce)
  fmt.Println(error)
  ---
  
  func functionError() (string, error) {
  	return "This is a string", errors.New("This is a new error")
  }
  ```
* The package `errors` implements functions to manipulate errors

## Panic
* While executing  a function F, an explicit call to the build in function `panic` or a `runtime panic` terminates the execution of F
* Any function deferred by F are executed as usual
* When all deferred functions are executed the program is terminated, and the error condition is reported, including the value of the argument to `panic`
* the termination sequence is called `panicking`

## Recover
* The recover function allows a program to manage behavior of a panicking goroutine
* The recover process:
  * Suppose a function G defers a function D that calls recover and a panic occurs in a function on the same goroutine in which G is executing. 
  * When the running of deferred functions reaches D, the return value of D's call to recover will be the value passed to the call of panic. 
  * If D returns normally, without starting a new panic, the panicking sequence stops 
  * In that case, the state of functions called between G and the call to panic is discarded, and normal execution resumes
  * Any functions deferred by G before D are then run and G's execution terminates by returning to its caller
