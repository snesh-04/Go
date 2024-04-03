package main

import "fmt"

func getMinMax(numbers []int) (int, int) {
    if len(numbers) == 0 {
        return 0, 0
    }

    min, max := numbers[0], numbers[0]
    for _, num := range numbers {
        if num < min {
            min = num
        }
        if num > max {
            max = num
        }
    }

    return min, max
}

func main() {
    numbers := []int{3, 7, 2, 8, 1, 9, 5}

    min, max := getMinMax(numbers)
    fmt.Printf("Minimum: %d\n", min)
    fmt.Printf("Maximum: %d\n", max)
}
// Write a program in GO language to demonstrate function return
// multiple values.