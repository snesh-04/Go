// Write a program in Go language how to create a channel and
// illustrate how to close a channel using for range loop and close
// function.
package main

import "fmt"

func main() {
    ch := make(chan int)

    go func() {
        defer close(ch)
        for i := 0; i < 5; i++ {
            ch <- i
        }
    }()

    for value := range ch {
        fmt.Println("Received:", value)
    }
}
