# Types
* **Declare**  a **variable**  of a certain **type**, it can only hold **values**  of that **type**
* Go is a  **static programming language**

## Classification data types
* Primitive data types
    * Basic type
    * Built-in type
* Composite data types 
    * Structure type
    * Aggregated type
    
## Boolean types
A boolean type represents a set of Boolean truth values denoted by the predeclared constants `true` and `false`

It allows compare variables with some expressions using operators. 

## Numeric types
You must use the specific numeric type to save memory space and time of processing, because your application can be leverage to millions of users, and each user at runtime will execute that function. (e.g. `uint8`, `uint16`, `int8`, `int16`, `float32`, `float64`, `complex64`, `complex128` )

## String types
A string type  represents the set of string values. A string value is a (possibly empty) sequence of bytes. String are immutable: once created, it is impossible to change the contents of a string.

( **\`** ) `back-text` allows you to put any character in your string including break lines and double quotes
    
 ## Own types
 * `type custom_type int`: declares a type named custom_type that will hold an integer. 
    * `int` is called **underlying type** 
    * The underlying type of everything is `T`
    * You can't assign a variable of type `int` to a variable of type `custom_type` because they are of different types
 * You can convert from type `int` to `custom_type` using `T(expression)`
   
# Values
## Zero value
* `bool`: false

* `string`: ""

* `int`: 0

* `float`: 0.0

* others: `nil`
    *  pointers
    * functions
    * interfaces
    * slices
    * channels
    * maps
    
 # Constants
You assign values at declaration time, once declared, it's value can't be changed

Constants can be TYPED or UNTYPED, an untyped constant can be implicitly converted by the compiler. 

If you assign the value iota you can make constants incremental (successive untyped integer constants),  it is reset to 0 whenever the reserved word const appears in the source. 