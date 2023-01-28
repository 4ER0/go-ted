# Goroutines
Go's `goroutines` are a way to run multiple functions concurrently within a single program. They are lightweight and efficient, making them a popular choice for concurrent programming in Go.

Here is an example of using goroutines in Go:

```golang
package main

import (
    "fmt"
    "time"
)

func printNumbers() {
    for i := 1; i <= 5; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}

func printLetters() {
    for i := 'a'; i < 'f'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}

func main() {
    go printNumbers()
    go printLetters()
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("\nmain function terminated")
}
```
In this example, we have two functions, printNumbers and printLetters, which are run concurrently using the go keyword. The main function runs both `printNumbers` and `printLetters` as `goroutines` and then sleeps for 3 seconds before printing "main function terminated" and exiting.

The output of this program will be interleaved numbers and letters, because the goroutines are running concurrently, the order of the output will vary each time the program is run.

```
1 a 2 b 3 c 4 d 5 
main function terminated
```
Note that the `time.Sleep` calls in the example are used to simulate some work being done, but you can use goroutines to run any function concurrently.