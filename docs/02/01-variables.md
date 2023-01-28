# Variables in Go

In Go, variables are used to store values that can be used and manipulated throughout the program. Variables can be declared and assigned a value in a single statement or in separate statements.

## Variable Declaration

A variable can be declared by using the keyword `var` followed by the variable name and its type. The type can be explicitly defined or inferred by the compiler.

```golang
var variableName type
```

For example:
```golang
var x int
var y string
var z float64
```

A variable can also be declared and assigned a value in the same statement using the shorthand notation `:=`.
```golang
variableName := value
```
For example:
```golang
x := 10
y := "Hello"
z := 3.14
```
It's also possible to declare multiple variables of the same type in one statement. 
```golang
var a, b, c int
```
## Variable Scope

The scope of a variable defines the portion of the program where the variable can be accessed. Variables declared inside a function are only accessible within that function, while variables declared outside of a function are accessible throughout the entire program.
```golang
package main

var globalVariable = "I'm accessible throughout the entire program"

func main() {
    localVariable := "I'm only accessible within the main function"
    fmt.Println(globalVariable)
    fmt.Println(localVariable)
}
```
## Constant

A constant is a value that cannot be changed after it is assigned. Constants are declared using the keyword `const` and can be of any data type.
```golang
const PI = 3.14
const MESSAGE = "Hello, World!"
```
It's also possible to declare multiple constants in one statement.
```golang
const (
    A = 1
    B = 2
    C = 3
)
```
In Go, it's best practice to use uppercase letters for constants.
This is for their easy identification and differentiation from variables in the source code.
