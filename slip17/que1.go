// Write a program in GO language to illustrate the concept of
// returning multiple values from a function. ( Add, Subtract,
// Multiply, Divide)
package main

import "fmt"

func arithmeticOperations(a, b float64) (float64, float64, float64, float64) {
    return a + b, a - b, a * b, a / b
}

func main() {
    num1 := 10.0
    num2 := 5.0

    sum, difference, product, quotient := arithmeticOperations(num1, num2)
    fmt.Printf("Sum: %.2f\n", sum)
    fmt.Printf("Difference: %.2f\n", difference)
    fmt.Printf("Product: %.2f\n", product)
    fmt.Printf("Quotient: %.2f\n", quotient)
}
