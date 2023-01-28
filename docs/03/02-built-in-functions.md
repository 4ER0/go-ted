# Built-in Functions
## make
`make` is a built-in function in Go that is used to create dynamic data structures such as arrays, slices, maps and channels. It requires a data type and an optional capacity as arguments.

Example:

```golang
// Create a slice with a capacity of 3
s := make([]int, 0, 3)
```
## append
`append` is a built-in function in Go that is used to add elements to the end of a slice. It takes a slice and one or more elements as arguments and returns a new slice with the additional elements.

Example:

```golang
s := []int{1, 2, 3}
s = append(s, 4, 5, 6)
// s is now [1, 2, 3, 4, 5, 6]
```
## len
`len` is a built-in function in Go that is used to determine the number of elements in a collection such as a slice, array, or string. It takes a collection as an argument and returns an integer.

Example:

```golang
s := []int{1, 2, 3}
l := len(s)
// l is now 3
```
## cap
`cap` is a built-in function in Go that is used to determine the capacity of a slice. It takes a slice as an argument and returns an integer.

Example:

```golang
s := make([]int, 0, 3)
c := cap(s)
// c is now 3
```
## new
`new` is a built-in function in Go that is used to create a new value of a given type and return a pointer to that value. It takes a type as an argument and returns a pointer to a new zero-value of that type.

Example:

```golang
p := new(int)
```
## delete
`delete` is a built-in function in Go that is used to delete key-value pairs from a map. It takes a map and a key as arguments.

Example:

```golang
m := map[string]int{"a": 1, "b": 2}
delete(m, "a")
// m is now map[b:2]
```
## copy
`copy` is a built-in function in Go that is used to copy elements from one slice to another. It takes two slices as arguments and returns the number of elements copied.

Example:

```golang
s1 := []int{1, 2, 3}
s2 := make([]int, 3)
n := copy(s2, s1)
// s2 is now [1, 2, 3] and n is 3
```
## close
`close` is a built-in function in Go that is used to close a channel. It takes a channel as an argument.

Example:

```golang
c := make(chan int)
close(c)
```
Please note that if you are using a channel after it has been closed, it will cause a panic.