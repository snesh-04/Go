// Write a program in GO language to swap two numbers using call
// by reference concept
package main

import "fmt"

func swapByReference(a *int, b *int) {
    temp := *a
    *a = *b
    *b = temp
}

func main() {
    var num1, num2 int
    fmt.Print("Enter first number: ")
    fmt.Scan(&num1)
    fmt.Print("Enter second number: ")
    fmt.Scan(&num2)

    fmt.Printf("Before swapping, num1 = %d and num2 = %d\n", num1, num2)
    swapByReference(&num1, &num2)
    fmt.Printf("After swapping, num1 = %d and num2 = %d\n", num1, num2)
}
