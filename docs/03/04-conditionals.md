# Conditionals

Go offers flow control with the conditional statements `if` and `switch`.

## If

Go's `if` statements are like its `for` loops; the expression need not be surrounded by parentheses `( )` but the braces `{ }` are required. 

For example ([Run code](https://play.golang.org/p/oxiKpLPSidg)):

```go
package main

import (
	"fmt"
)

func main() {
	bob := "PRODYNA"

	if bob == "PRODYNA" {
		fmt.Println("Hi " + bob + " colleague!")
	} else {
		fmt.Println("Hi stranger!")
	}
}
```

## If with a short statement

The `if` statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the if.

For example ([Run code](https://play.golang.org/p/Za8qekglXIj)):

```golang
package main

import (
	"fmt"
	"math"
)

func isPositive(n float64) bool {
	return n > 0
}

func main() {
	num := -10.2

	if v := math.Abs(num); isPositive(v) == true {
		fmt.Println("The number is positive")
	} else {
		fmt.Println("The number is negative")
	}
}
```

# Switch

A switch statement is a shorter way to write a sequence of `if - else` statements. It runs the first case whose value is equal to the condition expression.

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

For example ([Run code](https://play.golang.org/p/Ym-ZVpHxh9Y)):

```golang
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```

Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

Switch statement can accept variables as values. For example ([Run code](https://play.golang.org/p/0bqrdbGQTKl)):

```golang
package main

import (
	"fmt"
	"runtime"
)

var linuxVar = "linux"

func main() {

	switcher(linuxVar)
}

// switcher is just a function which performs
// just a switch statement and prints accordingly
func switcher(linuxArchitecture string) {

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case linuxArchitecture:
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```