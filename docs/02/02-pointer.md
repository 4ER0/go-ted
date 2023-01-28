# Pointers in Go

In Go, pointers are used to hold the memory address of a variable. They are declared using the `*` operator followed by the type of the variable that the pointer will point to.

## Pointer Declaration

A pointer variable can be declared by using the `*` operator followed by the variable name and the type of variable that it will point to.
```golang
var variableName *type
```
For example:
```golang
var x *int
var y *string
var z *float64
```
## Pointer Assignment

A pointer variable can be assigned the address of a variable using the `&` operator, which returns the memory address of a variable.
```golang
x := 10
var p *int = &x
```
It's also possible to use the `new` function to create a new variable and return its memory address.
```golang
p := new(int)
```
## Dereferencing Pointers

To access the value of the variable a pointer points to, the `*` operator is used. This is known as dereferencing the pointer.
```golang
fmt.Println(*p) // Output: 10
```
You can also use the pointer to modify the value of the variable it points to.
```golang
*p = 20
fmt.Println(x) // Output: 20
```

# Memory Model

## Values
* Values are normally stored in the stack and therefore very fast.
* Values can be still modified. But it results in a copy.
* A value is not garbage collected because it will be cleared when the stack frame is destroyed 

## Pointers
* A Pointer references a value that is stored in the heap
* When pointers are not longer used, they need to be garbage collected
