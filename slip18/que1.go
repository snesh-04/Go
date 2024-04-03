// Write a program in GO language to print a multiplication table of
// number using function.
package main

import "fmt"

func multiplicationTable(num int) {
    for i := 1; i <= 10; i++ {
        fmt.Printf("%d x %d = %d\n", num, i, num*i)
    }
}

func main() {
    number := 7
    multiplicationTable(number)
}
