# Functions
Functions is about modularize code

* We define our **function** with **parameters** (declaration)

* We call our **function** and pass in **arguments** (invocation)

**Everything** in go **is passed by value**


* Declaration
  ```go
  func (r recevier) indentifier(parameters) (return(s)) { 
   ... 
  }
  ```
* Invocation
  ```go
  return(s) := identifier(argument(s))
  ```

## Variadic Parameters
You can pass **zero to many arguments** into single one **variadic parameter**, the variadic parameter will create a slice with all the specified values

* Declaration
    ```go
    func variadic(x ...int)  {
     ...
    }
    ```
* Invocation
    ```go
    variadic(value_1,value_2,...,value_n)
    ```

## Defer statements
A `defer` statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because corresponding goroutine is panicking. **That is, the `defer` function will be executed when the surrounding function returns or exits**.
```go
func surroundingFunction()  {
	defer deferFunction() 
    ...
    code
} // deferFunction will be executed here
```

## Methods
Methods are attached to a struct through a receiver
```go
type structName struct {
	field field_type
}

// func (r receiver)
func (s structName) indentifier(parameters) (return(s)) { 
 ...
}
```
### Method Sets
* Method Sets: [Pointers](../08_pointer#method-sets) (Advanced topic, see it after Interfaces and Pointers)

## Interfaces
* Interfaces allow things to work well together to interface together.
* A value can be of more than one type. 
* Any type that implements all the interface's methods is also of the interface's type.
* Declaration
  ```go
  type interfaceName interface {
  	methodInterface()
  }
  ```
* Method implementation
  ```go
  func (p structName1) methodInterface()  {
  }
  
  func (p structName2) methodInterface()  {
  }
  ```
* Usage
  ```go
  func anotherFunction(i interfaceName)  {
  	i.methodInterface()
  }
  ```
* Invocation
  ```go
  s1 := structName1{}
  s2 := structName2{}
  
  anotherFunction(s1) // Same function, different type,
  anotherFunction(s2) // But same interface
  ```
* Assertion
  ```go
  i.(structName1).field
  ```
* Every type implement the interface `interface{}` because it has no methods because all types implements at least no methods

## Anonymous 
* Definition and invocation
  ```go
  func(parameter parameter_type){ 
    code
  }(argument) 
  ```
  
## Expression
Functions can be used just like any other variable or any other type. 
*'Function is a type'*
* Definition
  ```go
  f := func(parameter parameter_type){
    code
  }
  ```
* Invocation
  ```go
  f(argument)
  ```
  
## Return a function
* Definition
  ```go
  func calledFunction() func() func_return_type {
    return func() func_return_type {
      code
    }
  }
  ```
* Invocation
  ```go
  function_returned := calledFunction()
  function_returned()
  ```
  
## Callback
* Definition
  ```go
  func callbackFunction(name_callback func(parameter_c type_c) return_c, parameter_1 type_1) return_1  {
    code
    ...
    return_value := name_callback(parameter_1)
    return return value
  }
  ```
* Invocation
  ```go
  f := callbackFunction(name_callback,parameter)
  ```
 ## Closure
Closure is about variable scope, the scope is defined by `{}`, it can be a function or just a code block.

* See the [the closure example](08_closure.go)

## Recursion
Preference loops over recursion, recursion is unnecessrily complex and it can have negative impact on memory usage.

* See the [the recursion example](09_recursion.go)
