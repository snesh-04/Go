// Write a program in GO language to illustrate the function
// returning multiple values(add, subtract).
package main

import "fmt"

func addAndSubtract(a, b int) (int, int) {
    sum := a + b
    diff := a - b
    return sum, diff
}

func main() {
    num1 := 20
    num2 := 10

    sum, diff := addAndSubtract(num1, num2)
    fmt.Printf("Sum: %d, Difference: %d\n", sum, diff)
}
