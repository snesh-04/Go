// Write a program in GO language to accept two matrices and
// display its multiplication
package main

import "fmt"

func multiplyMatrices(matrix1 [][]int, matrix2 [][]int) [][]int {
    rows1 := len(matrix1)
    cols1 := len(matrix1[0])
    cols2 := len(matrix2[0])

    result := make([][]int, rows1)
    for i := range result {
        result[i] = make([]int, cols2)
    }

    for i := 0; i < rows1; i++ {
        for j := 0; j < cols2; j++ {
            for k := 0; k < cols1; k++ {
                result[i][j] += matrix1[i][k] * matrix2[k][j]
            }
        }
    }

    return result
}

func main() {
    var rows1, cols1, rows2, cols2 int
    fmt.Print("Enter the number of rows of matrix 1: ")
    fmt.Scan(&rows1)
    fmt.Print("Enter the number of columns of matrix 1: ")
    fmt.Scan(&cols1)

    fmt.Print("Enter the number of rows of matrix 2: ")
    fmt.Scan(&rows2)
    fmt.Print("Enter the number of columns of matrix 2: ")
    fmt.Scan(&cols2)

    if cols1 != rows2 {
        fmt.Println("Matrix multiplication not possible.")
        return
    }

    matrix1 := make([][]int, rows1)
    fmt.Println("Enter elements of matrix 1:")
    for i := 0; i < rows1; i++ {
        matrix1[i] = make([]int, cols1)
        for j := 0; j < cols1; j++ {
            fmt.Printf("Enter element [%d][%d]: ", i, j)
            fmt.Scan(&matrix1[i][j])
        }
    }

    matrix2 := make([][]int, rows2)
    fmt.Println("Enter elements of matrix 2:")
    for i := 0; i < rows2; i++ {
        matrix2[i] = make([]int, cols2)
        for j := 0; j < cols2; j++ {
            fmt.Printf("Enter element [%d][%d]: ", i, j)
            fmt.Scan(&matrix2[i][j])
        }
    }

    result := multiplyMatrices(matrix1, matrix2)

    fmt.Println("Result of matrix multiplication:")
    for _, row := range result {
        fmt.Println(row)
    }
}
