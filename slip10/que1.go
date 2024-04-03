// Write a program in GO language to create an interface and display
// its values with the help of type assertion.
package main

import (
    "fmt"
)

type shape interface {
    area() float64
}

type rectangle struct {
    length, width float64
}

func (r rectangle) area() float64 {
    return r.length * r.width
}

type circle struct {
    radius float64
}

func (c circle) area() float64 {
    return 3.14 * c.radius * c.radius
}

func printArea(s shape) {
    switch v := s.(type) {
    case rectangle:
        fmt.Printf("Area of rectangle: %.2f\n", v.area())
    case circle:
        fmt.Printf("Area of circle: %.2f\n", v.area())
    }
}

func main() {
    r := rectangle{length: 5, width: 3}
    c := circle{radius: 5}

    printArea(r)
    printArea(c)
}
