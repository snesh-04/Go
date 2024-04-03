// B)
// Write a program in GO language to copy all elements of one array
// into another using a method.
package main

import "fmt"

func copyArray(source []int, destination []int) {
    for i := 0; i < len(source); i++ {
        destination[i] = source[i]
    }
}

func main() {
    source := []int{1, 2, 3, 4, 5}
    destination := make([]int, len(source))

    copyArray(source, destination)

    fmt.Println("Source array:", source)
    fmt.Println("Destination array:", destination)
}
