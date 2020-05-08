# Structs
A struct is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified **explicitly (IdentifierList)** or **implicitly (EmbeddedField).** Within a struct, field names must be unique.

A struct is a composite data type. Structs allow us to compose together values of different types.

## Identifier List
* Struct declaration
  ```go
  type struct_name struct {
    field_1 type_field_1
    field_2 type_field_2
    ...
    field_n type_field_3
  }
  ```
* Variable declaration
  ```go
  p1 := struct_name{
    field_1: "value_1",
    field_2: "value_2",
  }
  ```
* Accessing fields
  ```go
  x := p1.field_1
  ```
## Embedded Field
A field declared with a type but no explicit field name is called an embedded field. The unqualified type name acts as the field name.
* Struct embedded declaration
  ```go
  type embedded_struct struct {
    field_embedded_1 type_field_1
    field_embedded_2 type_field_2
    ...
    field_embedded_n type_field_3
  }
  
  type composed_struct struct {
    embedded_struct // The previous struct 
    
    field_1 type_field_1
    field_2 type_field_2
    ...
    field_n type_field_3
  }
  ```
* Variable declaration
  ```go
  p1 := composed_struct {  // `composed_struct` type
    embedded_struct: embedded_struct {  // `embedded_struct` type
      field_embedded_1: "value_1",
      field_embedded_2: "value_2",
    }, 
    field_1: "value_1",
  }
  ```
* Accessing fields 
  ```go
  x := p1.field_embedded_1 // Field promotion
  x := p1.embedded_struct.field_embedded_1 // It's also valid (use when name collision)
  ```
  
# Anonymous Structs
It's anonymous because it doesn't have name
```go
p1 := struct {  // Anonymous struct
        field_1 type_1
        field_2 type_2
      }{
        field_1: value_1,
        field_2: value_2,
      }
```