// Write a program in GO language to create a user defined package
// to find out the area of a rectangle.
package main

import (
    "fmt"
    "geometry"
)

func main() {
    length := 5.0
    width := 3.0

    area := geometry.AreaOfRectangle(length, width)
    fmt.Printf("Area of rectangle with length %.2f and width %.2f is %.2f\n", length, width, area)
}
