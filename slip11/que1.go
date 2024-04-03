// Write a program in GO language to check whether the accepted
// number is two digit or not.
package main

import "fmt"

func isTwoDigit(num int) bool {
    return num >= 10 && num <= 99
}

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if isTwoDigit(num) {
        fmt.Println(num, "is a two-digit number.")
    } else {
        fmt.Println(num, "is not a two-digit number.")
    }
}
