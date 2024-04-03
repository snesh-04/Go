// Write a program in GO language using a function to check
// whether the accepted number is palindrome or not.
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
        num /= 10
    }

    return originalNum == reversedNum
}

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if isPalindrome(num) {
        fmt.Println(num, "is a palindrome.")
    } else {
        fmt.Println(num, "is not a palindrome.")
    }
}
