// A) Write a program in GO language to print Fibonacci series of n
// terms.

package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 int
	var n, count int
	count = 2
	fmt.Println("Enter how many terms want???")
	fmt.Scanln(&n)

	num1 = 0
	num2 = 1
	fmt.Print("\n", num1, ",", num2)

	for count = 2; count < n; count++ {
		num3 = num1 + num2
		fmt.Print(",", num3)
		num1 = num2
		num2 = num3
	}
}
