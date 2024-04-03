// B) Write a program in GO language that creates a slice of integers,
// checks numbers from the slice are even or odd and further sent to
// respective go routines through channel and display values
// // received by goroutines.
package main

import (
    "fmt"
)

func checkEvenOdd(num int, evenCh chan<- int, oddCh chan<- int) {
    if num%2 == 0 {
        evenCh <- num
    } else {
        oddCh <- num
    }
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    evenCh := make(chan int)
    oddCh := make(chan int)

    for _, num := range nums {
        go checkEvenOdd(num, evenCh, oddCh)
    }

    for i := 0; i < len(nums); i++ {
        select {
        case evenNum := <-evenCh:
            fmt.Printf("Received even number: %d\n", evenNum)
        case oddNum := <-oddCh:
            fmt.Printf("Received odd number: %d\n", oddNum)
        }
    }
}
