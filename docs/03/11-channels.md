# Channels
In Go, channels are a way to communicate between goroutines (lightweight threads) and synchronize their execution. A channel is a type that can be used to send and receive values of a certain type. Channels are created using the make function, which takes the type of the values that the channel will hold as an argument.

Here is an example of a simple use of channels:

```golang
package main

import "fmt"

func sum(a, b int, c chan int) {
    c <- a + b
}

func main() {
    c := make(chan int)
    go sum(1, 2, c)
    result := <-c
    fmt.Println(result) // output: 3
}
```
In this example, we create a channel c of type int using make(chan int). We then use the go keyword to run the sum function in a separate goroutine, passing the channel as an argument. The sum function sends the result of the addition to the channel using the <- operator. In the main function, we use the <- operator again to receive the value from the channel and print it.

## Buffered Channels
By default, channels are unbuffered, meaning that a sender will block until there is a receiver ready to receive the value. Buffered channels allow a certain number of values to be sent without blocking. They are created using the make function with an additional argument indicating the buffer size.

```golang
c := make(chan int, 2)
```
If the buffer is full, the sender will block until a receiver receives a value. If the buffer is empty, the receiver will block until a value is sent.

## Closing Channels
A channel can be closed using the close function. Once a channel is closed, any further sends will result in a panic, and any further receives will return the zero value of the type.

```golang
package main

import "fmt"

func main() {
    c := make(chan int)
    go func() {
        c <- 1
        c <- 2
        close(c)
    }()

    for value := range c {
        fmt.Println(value)
    }
    // Output: 1 2
}
```
In this example, we create a channel c and run a goroutine that sends two values to the channel and then closes it. The for loop in the main function receives the values from the channel using the range keyword, which automatically stops when the channel is closed.