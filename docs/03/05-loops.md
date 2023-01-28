# Loops

## For

Go has only one looping construct, the `for` loop.

The basic `for` loop has three components separated by semicolons:

- the init statement: executed before the first iteration
- the condition expression: evaluated before every iteration
- the post statement: executed at the end of every iteration

The init statement will often be a short variable declaration, and the
variables declared there are visible only in the scope of the `for`
statement.

The loop will stop iterating once the boolean condition evaluates to `false`.

For example ([Run code](https://play.golang.org/p/Yh8jRtIdbuT)):
```golang
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

**Note:** Unlike other languages like C, Java, or JavaScript there are no parentheses
surrounding the three components of the `for` statement and the braces `{}` are
always required.

The init and post statements are *optional*.

For example ([Run code](https://play.golang.org/p/iLbpPmG0di5)):

```golang
package main

import "fmt"

func main() {
	sum := 1
    // Check the fact that we do not initialise
    // the sum inside for loop.
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
```

## Where is "while" or "Do-while"?

There is no `while` or `do - while`. Go has only `for` which can act like `while` or `do - while`.

For example ([Run code](https://play.golang.org/p/-8ihiw0pxlX)):

```golang
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```

## Infinite loop

If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

```golang
package main

func main() {
	for {
	}
}
```

##  Range

The `range` form of the `for` loop iterates over a slice or map.

### Range over a slice

When ranging over a slice, two values are returned for each iteration. The first is the `index`, and the second is a copy of the element at that index.

For example ([Run code](https://play.golang.org/p/MoSqZXcLtkl)):
```go
package main

import "fmt"

var names = []string{"Frank", "Martin", "Chris"}

func main() {
	for i, v := range names {
		fmt.Printf("%d. %s writes Go!\n", (i + 1), v)
	}

	// Print the names again but omit the index
	for _, v := range names {
		fmt.Printf("%s writes Go!\n", v)
	}
}
```

### Range over a map

When ranging over a map, two values are returned for each iteration. The first is the `key` of the map and the second is the `value` of that `key`.

For example ([Run code](https://play.golang.org/p/TAiIwLs3hEJ)):
```golang
package main

import "fmt"

func main() {
	var machMap = map[string]float64{
		"F-16":   1.6,
		"SR-71":  3.2,
		"Mig-25": 2.83,
	}

	fmt.Println(machMap)

	for k, v := range machMap {
		fmt.Printf("%s can go up to mach: %.2f!\n", k, v)
	}

	// Print the keys again but omit the value
	fmt.Println("\nPrint only airplane names:")
	for k, _ := range machMap {
		fmt.Println(k)
	}
}
```