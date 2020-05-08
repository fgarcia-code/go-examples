# Pointers
* Every **value** in any computer is **stored** in **memory**

* To get **access** to any **value**, the variables must **reference** to a **memory location**, this is called the **memory address**

* Every **variable** has a **memory address** that **points** to a **value**, you can get the memory address using the `&` operator:
  ```go
  // Declaring a variable
  variable := value

  // Accessing the value of the variable
  fmt.Println(variable)

  // Getting the memory address of variable
  fmt.Println(&variable)
  ```
* The **type** of the **value's memory address** is of type `*type` (where `type` is the **type's value**)

* You can create a **variable of type** `*type`, using the short variable declaration
  ```go
  // pointer_variable is of type `*type`
  pointer_variable := &variable
  ```
* The **value** of `pointer_variable` is **a memory address**
* You can **access** the **value** that is **pointed by** `pointer_variable` using the `*` **operator** (**the value** will be the same as the `variable` value)
  ```go
  // When you use `*` operator you are dereferencing the memory address
  fmt.Println(*pointer_variable)
  ```
* If the `variable` value **changes**, the `*pointer_variable` will also change (it is the **same address**)

**Note**: Everything in go is pass by value, memory address is also pass by value

## When to use pointers
* Pointers are good if you have a large chunk of data, and you don't want to pass that big chunk of data around though your program, you just pass that address where that data is stored, this helps to save you a little performance

* Pointers ar also good if you want to change (mutate) some value stored at some specific memory address

## Method Sets
* A **type** may **have a method set** associated with it
* The **method set** of an **interface** type **is its interface** (Method's interface)
* *The **method set** of **type T** consists of all **methods declared** with **receiver type T***
* *The **method set** of the corresponding **pointer type \*T** is the **set of all methods** declared with **receiver \*T or T***
* Any other type has an empty method set
* The **method set** of a type **determines** the **interfaces** that the **type implements** and the **methods** that **can be called** using a **receiver** of that **type**
