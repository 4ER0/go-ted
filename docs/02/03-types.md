# Data Types in Go

Go has several built-in data types, including:

## Numbers

- Integers: int, uint, int8, int16, int32, int64, uint8, uint16, uint32, uint64
- Floating-point: float32, float64
- Complex numbers: complex64, complex128

## Strings

- String: a sequence of Unicode characters

## Booleans

- bool: true or false

## Arrays

- Arrays have a fixed length and can hold multiple values of the same type
- Example: 
```golang
    var a [5]int
```
## Slices

- Slices are a more flexible version of arrays and are more commonly used in Go
- They can be resized and have a dynamic length
- Example: 
```golang
    var b []int
```
## Maps

- Maps are a collection of key-value pairs
- Example:
```golang
    var m map[string]int
```
## Structs

- Structs are a composite data type that can hold multiple fields
- Example:
```golang
    type User struct {
        Name string
        Age int
    }
```