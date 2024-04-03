// Write a function in GO language to find the square of a number
// and write a benchmark for it.
package main

import "testing"

func square(num int) int {
    return num * num
}

func BenchmarkSquare(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = square(5) // Change the number here for benchmarking different values
    }
}
