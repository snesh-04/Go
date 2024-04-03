// B) Write a program in GO language to create a buffered channel,
// store few values in it and find channel capacity and length. Read
// values from channel and find modified length of a channel
package main

import "fmt"

func main() {
    ch := make(chan int, 3)

    // Store values in the buffered channel
    ch <- 1
    ch <- 2
    ch <- 3

    // Find channel capacity and length
    fmt.Println("Channel capacity:", cap(ch))
    fmt.Println("Initial channel length:", len(ch))

    // Read values from channel
    fmt.Println("Values from channel:")
    for i := 0; i < 3; i++ {
        fmt.Println(<-ch)
    }

    // Find modified length of channel
    fmt.Println("Final channel length:", len(ch))
}
