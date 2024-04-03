// Write a program in GO language that prints out the numbers from0 
// to 10, waiting between 0 and 250 ms after each one using the
// delay function
package main

import (
    "fmt"
    "time"
)

func delayPrintNumbers() {
    for i := 0; i <= 10; i++ {
        fmt.Println(i)
        time.Sleep(time.Duration(i*25) * time.Millisecond) // Wait between 0 and 250 ms
    }
}

func main() {
    delayPrintNumbers()
}
