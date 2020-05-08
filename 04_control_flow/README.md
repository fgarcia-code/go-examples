# Control Flow
## If statement (Conditional)
"If" statements specify the conditional execution of two branches according to the value of a boolean expression. If the expression evaluates to true, the "if" branch is executed, otherwise, if present, the "else" branch is executed.
* Single condition
    ```go
    if expression {}
    ```
* If with init statement
    ```go
    if init; expression {}
    ```
* If else if
    ```go
    if expression {} 
    else if expression {} 
    else {}
    ```
## Switch statement (Conditional)
"Switch" statements provide multi-way execution. An expression or type specifier is compared to the "cases" inside the "switch" to determine which branch to execute. 
### Expression switches 
In an expression switch, the cases contain expressions that are compared against the value of the switch expression. The switch expression is evaluated exactly once in a switch statement.

 The switch expression is evaluated and the case expressions, which need not be constants, are evaluated left-to-right and top-to-bottom; the first one that equals the switch expression triggers execution of the statements of the associated case; the other cases are skipped.

* Expression switch comparison
    ```go
    switch {
      case expression:
        f()
        fallthrough
      case expression:
        f()
        fallthrough  
      default:
        f()
    }
    ```
* Expression switch tag
    ```go
    switch tag {
      case option_1, option_2, option_3, option_4:
        f()
      case option_4, option_5, option_6, option_4:
        f()
    }
    ```
* Expression switch tag
    ```go
    switch option := expression_function(); {
      case expression:
        f()
      case expression:
        f()
    }
    ```
### Type switches
In a type switch, the cases contain types that are compared against the type of a specially annotated switch expression. The switch expression is evaluated exactly once in a switch statement.
* A type switch compares types rather than values. It is otherwise similar to an expression switch.
    ```go
    switch x.(type) {	 
      case int:
        f()	 
    }
    ```
## For statement (Loop)
A "for" statement specifies repeated execution of a block. There are three forms: The iteration can be controlled by a single condition, a "for" clause, or a "range" clause.

* Empty for
    ```go
    for {}
    ```
* Single condition (while)
    ```go
    for condition {}
    ```    
* For clause
    ```go
    for init; condition; post {}
    ```    
* Range clause
    ```go
    for expression_list := range expression {}
    ```