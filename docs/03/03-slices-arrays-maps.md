# Arrays, Slices, and Maps in Go

Go provides several data types for storing and working with collections of data. This includes arrays, slices, and maps.

## Arrays

An array is a fixed-size collection of elements of the same type. The size of an array is set when it is created and cannot be changed afterwards. Here is an example of creating an array of integers:

```golang
var a [3]int
a[0] = 1
a[1] = 2
a[2] = 3
```
You can also initialize an array at the time of creation by providing a list of values enclosed in curly braces:

```golang
a := [3]int{1, 2, 3}
```

### Multidimensional Arrays
Go supports multidimensional arrays, which can be created by specifying multiple dimensions in the square brackets. Here is an example of creating a 2-dimensional array of integers:

```golang
var b [2][3]int
b[0][0] = 1
b[0][1] = 2
b[0][2] = 3
b[1][0] = 4
b[1][1] = 5
b[1][2] = 6

c := [2][3]int{{1, 2, 3}, {4, 5, 6}}
```
This creates an array with 2 rows and 3 columns. Elements in a multidimensional array can be accessed using multiple indices, like b[i][j]

## Slices
A slice is a flexible-size collection of elements of the same type. The size of a slice can be changed after it is created. Here is an example of creating a slice of integers:

```golang
s := []int{1, 2, 3}
```
Slices can also be created using the make function, which allows you to specify the initial capacity and length:

```golang
s := make([]int, 0, 3)
```

You can access and modify the elements of a slice in the same way as an array:

```golang
fmt.Println(s[1]) // prints 1
s[0] = 100
```
To add a new element to that slice you need to use the append function

```golang
s = append(s, 4)
```
## Maps
A map is a collection of key-value pairs, where each key is unique. Here is an example of creating a map of strings to integers:

```golang
m := map[string]int{
    "one": 1,
    "two": 2,
    "three": 3,
}
```

You can also create a map using the make function:

```golang
m := make(map[string]int)
m["one"] = 1
m["two"] = 2
m["three"] = 3
```
To add or update a value in a map, use the following syntax:

```golang
m["four"] = 4
```

You can use the delete function to remove an element from a map:

```golang
delete(m, "one")
```
## Working with Arrays, Slices, and Maps
All of these data types have built-in functions and methods for working with them. For example, you can use the len function to get the number of elements in an array or slice and the range keyword to iterate over the elements of an array, slice, or map.

```golang
for k, v := range m {
    fmt.Printf("%s: %d\n", k, v) //%s means that the first var 'k' is an string while %d means 'v' is and base 10 integer
}
```

It's important to note that maps, slices and arrays are reference types in Go and when passed as argument in function or method it is passed as reference, so any changes made inside function will be reflected in the original data structure.

These data types are fundamental building blocks for many Go programs and are widely used in different scenarios.