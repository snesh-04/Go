// Write a program in GO language to sort array elements in
// ascending order.
package main

import (
    "fmt"
    "sort"
)

func main() {
    arr := []int{64, 34, 25, 12, 22, 11, 90}
    fmt.Println("Unsorted array:", arr)
    sort.Ints(arr)
    fmt.Println("Sorted array:", arr)
}
