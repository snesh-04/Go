// Write a program in GO language to print a recursive sum of digits
// of a given number.
package main

import "fmt"

func recursiveSumOfDigits(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + recursiveSumOfDigits(n/10)
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	sum := recursiveSumOfDigits(num)
	fmt.Printf("Recursive sum of digits of %d is %d\n", num, sum)
}
