// Write a program in GO language to print sum of all even and odd
// numbers separately between 1 to 100.
package main

import "fmt"

func main() {
    sumEven, sumOdd := 0, 0

    for i := 1; i <= 100; i++ {
        if i%2 == 0 {
            sumEven += i
        } else {
            sumOdd += i
        }
    }

    fmt.Println("Sum of even numbers between 1 to 100:", sumEven)
    fmt.Println("Sum of odd numbers between 1 to 100:", sumOdd)
}
