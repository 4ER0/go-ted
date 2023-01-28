# Functions
In Go, functions are defined using the func keyword, followed by the function name, a list of parameters, and a block of code to be executed when the function is called. Functions can also return one or more values.

## Scope

The scope of a function is determined by the first letter
* Uppercase \
  The function is public and everywhere available
* Lowercase \
  The function is private/package local and can be called only within the module/package


## Function Syntax
```golang
func functionName(parameter1 type1, parameter2 type2) (returnType1, returnType2) {
    // code to be executed
    return value1, value2
}
```

Here is an example of a simple function that takes two integers as parameters and returns their sum:

```golang
func add(a int, b int) int {
    return a + b
}

result := add(3, 4)
fmt.Println(result) // output: 7
```
##  Multiple return values
A function can return multiple values

```golang
func addAndMultiply(a int, b int) (int, int) {
    return a + b, a * b
}

sum, product := addAndMultiply(3, 4)
fmt.Println("sum:", sum) // output: 7
fmt.Println("product:", product) // output: 12
```
## Named return values
Go also allows you to give names to the return values of a function, which can make the code more readable.

```golang
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

a, b := split(7)
fmt.Println("a:", a) // output: 3
fmt.Println("b:", b) // output: 4
```
## Variadic function
A function can accept a variable number of arguments by using the ... syntax.

```golang
func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}

result := add(1, 2, 3, 4, 5)
fmt.Println(result) // output: 15
```

# Receiver Functions

In go, functions are not attached to anything. They can be called from everywhere unless they are public.
To attach functions to struct, the function needs a receiver.

```golang
type Point struct {
    X int
    Y int
}

func (p Point) GetX() int {
	return p.X
}
```

The Receiver Function comes in to variations
- Value Receiver
- Pointer Receiver

## Value Receiver
A Value Receiver Function cannot change the internal state of the struct.

```golang
type Point struct {
    X int
    Y int
}

func (p Point) Coordinates() (int, int) {
    return p.X, p.Y
}

```

## Pointer Receiver
A Point Receiver can change the internal state of the struct.


```golang
type Point struct {
    X int
    Y int
}

func (p *Point) SetX(x int) {
	p.X = x
}
```