// ) Write a program in GO language to demonstrate working of slices
// (like append, remove, copy etc.)
package main

import "fmt"

func main() {
    // Create a slice
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println("Original Slice:", slice)

    // Append elements to slice
    slice = append(slice, 6, 7, 8)
    fmt.Println("After Append:", slice)

    // Remove element at index 2
    indexToRemove := 2
    slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)
    fmt.Println("After Remove at Index 2:", slice)

    // Copy slice to a new slice
    copiedSlice := make([]int, len(slice))
    copy(copiedSlice, slice)
    fmt.Println("Copied Slice:", copiedSlice)
}
