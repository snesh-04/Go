// Write a program in GO language to accept user choice and print answers
// using arithmetic operators.
package main

import (
	"fmt"
)

func main() {
	var choice string
	var num1, num2, result int

	fmt.Println("Enter your choice of operation (+, -, *, /):")
	fmt.Scanln(&choice)

	fmt.Println("Enter two numbers:")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)

	switch choice {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Cannot divide by zero!")
			return
		}
	default:
		fmt.Println("Invalid choice!")
		return
	}

	fmt.Printf("Result: %d\n", result)
}
