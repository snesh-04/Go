// Write a program in GO language using go routine and channel that
// will print the sum of the squares and cubes of the individual digits
// of a number. Example if number is 123 then
// squares = (1 * 1) + (2 * 2) + (3 * 3)
// cubes = (1 * 1 * 1) + (2 * 2 * 2) + (3 * 3 * 3)
package main

import (
    "fmt"
    "strconv"
)

func digitSquaresCubesSum(num int, sumSquaresCh chan int, sumCubesCh chan int) {
    digits := strconv.Itoa(num)
    sumSquares := 0
    sumCubes := 0

    for _, digitStr := range digits {
        digit := int(digitStr - '0')
        sumSquares += digit * digit
        sumCubes += digit * digit * digit
    }

    sumSquaresCh <- sumSquares
    sumCubesCh <- sumCubes
}

func main() {
    num := 123
    sumSquaresCh := make(chan int)
    sumCubesCh := make(chan int)

    go digitSquaresCubesSum(num, sumSquaresCh, sumCubesCh)

    sumSquares := <-sumSquaresCh
    sumCubes := <-sumCubesCh

    fmt.Printf("Number: %d\n", num)
    fmt.Printf("Sum of squares of digits: %d\n", sumSquares)
    fmt.Printf("Sum of cubes of digits: %d\n", sumCubes)
}
