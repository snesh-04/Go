// A) Write a program in the GO language using function to check
// whether accepts number is palindrome or not.
package main

import (
	"fmt"
)

func isPalindrome(num int) bool {
	originalNum := num
	reversedNum := 0

	for num > 0 {
		remainder := num % 10
		reversedNum = reversedNum*10 + remainder
		num = num / 10
	}

	return originalNum == reversedNum
}

func main() {
	var num int
	fmt.Println("Enter a number:")
	fmt.Scanln(&num)

	if isPalindrome(num) {
		fmt.Printf("%d is a palindrome\n", num)
	} else {
		fmt.Printf("%d is not a palindrome\n", num)
	}
}
