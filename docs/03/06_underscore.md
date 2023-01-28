# Underscore

In Go, the underscore `_` is used as a placeholder when you don't need or want to use a variable.

## Discarding Values
When you use the `_` in place of a variable, it discards the value that would have been assigned to that variable. This can be useful when you are iterating over a range and don't need the index or value, for example:

```golang
numbers := []int{1, 2, 3, 4, 5}
for _, number := range numbers {
    fmt.Println(number)
}
```
In this case, the index is being discarded and only the value is being printed.

## Ignoring Errors
Another common use of _ is to ignore errors that are returned by a function.

```golang
_, err := someFunction()
if err != nil {
    fmt.Println("An error occured")
}
```
In this case, the function returns two values, one of which is an error. By using _ in place of the first value, we are indicating that we don't need it and only care about the error.

## Blank Identifier
The underscore is also referred as a blank identifier that can be used to ignore values returned by function or values assigned to variables.
