# Grouping Data
## Array
An array is a numbered sequence of elements of a single type, called the element type. The number of elements is called the length of the array and is never negative. The length is part of the array's type.
```go
var x [size]int
```
## Slice
Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data, most array programming in Go is done with slices rather than simple arrays.

Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array.

It only store values of the same type.

* Slice declaration
  ```go
  x := []type{values}
  ```
* Iterate through the slice
  ```go
  for i, v := range `slice` { }
  ```
* Slicing the slice
  ```go
  x[from_index:until_index]
  ```
* Append to slice
  ```go
  x = append(x, element_1, element_2, element_3)
  x = append(x, other_slice...)
  ```
* Delete from slices (exclude the 3rd element with the append function)
  ```go
  x = append(x[:2],x[4:]...)
  ```
* Slice make (it creates a slice, the memory of the underlying array will be reserved in blocks of the capacity)
  ```go
  x := make([]T, length, capacity)
  ```
* Multi-dimensional slice
  ```go
  x := [][]string{{v1,v2},{v4,v5}}
  ```
  
## Map
A map is an unordered group of elements of one type, called the element type, indexed by a set of unique keys of another type, called the key type. The value of an uninitialized map is nil.

Like slices, maps hold references to an underlying data structure.

* Declaration
  ```go
  m := map[type_key]type_value{}
  ```
* Accessing an element
  ```go
  m[key]
  ```
  If you access an element that isn't present you will get the zero value
* You can check whether a key exist or not, in the following example the identifier `ok` is of type boolean
  ```go
  if _,ok := m[key]; !ok {
    f(x)
  }
  ```
* Adding an element
  ```go
  m[new_key] = value
  ```
* Iterating the map
  ```go
  for k, v := range m {
    f(k,v)
  }
  ```
* Deleting an element
  ```go
  delete(m,key)
  ```
