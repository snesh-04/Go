// Write a program in GO language using a user defined package
// calculator that performs one calculator operation as per the user's
// choice.
package main

import (
    "fmt"
    "calculator"
)

func main() {
    var choice int
    var num1, num2 float64

    fmt.Println("Select operation:")
    fmt.Println("1. Add")
    fmt.Println("2. Subtract")
    fmt.Println("3. Multiply")
    fmt.Println("4. Divide")
    fmt.Print("Enter choice (1-4): ")
    fmt.Scan(&choice)

    fmt.Print("Enter first number: ")
    fmt.Scan(&num1)
    fmt.Print("Enter second number: ")
    fmt.Scan(&num2)

    switch choice {
    case 1:
        result := calculator.Add(num1, num2)
        fmt.Printf("Result: %.2f\n", result)
    case 2:
        result := calculator.Subtract(num1, num2)
        fmt.Printf("Result: %.2f\n", result)
    case 3:
        result := calculator.Multiply(num1, num2)
        fmt.Printf("Result: %.2f\n", result)
    case 4:
        result := calculator.Divide(num1, num2)
        fmt.Printf("Result: %.2f\n", result)
    default:
        fmt.Println("Invalid choice")
    }
}
