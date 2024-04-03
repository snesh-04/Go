// Write a program in GO language to accept one matrix and display
// // its transpose.
package main

import "fmt"

func transpose(matrix [][]int) [][]int {
    rows := len(matrix)
    cols := len(matrix[0])

    result := make([][]int, cols)
    for i := range result {
        result[i] = make([]int, rows)
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            result[j][i] = matrix[i][j]
        }
    }

    return result
}

func main() {
    var rows, cols int
    fmt.Print("Enter the number of rows of the matrix: ")
    fmt.Scan(&rows)
    fmt.Print("Enter the number of columns of the matrix: ")
    fmt.Scan(&cols)

    matrix := make([][]int, rows)
    fmt.Println("Enter elements of the matrix:")
    for i := 0; i < rows; i++ {
        matrix[i] = make([]int, cols)
        for j := 0; j < cols; j++ {
            fmt.Printf("Enter element [%d][%d]: ", i, j)
            fmt.Scan(&matrix[i][j])
        }
    }

    result := transpose(matrix)

    fmt.Println("Transpose of the matrix:")
    for _, row := range result {
        fmt.Println(row)
    }
}
