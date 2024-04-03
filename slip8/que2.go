// B) Write a program in GO language to create an interface shape that
// includes area and perimeter. Implements these methods in circle
// and rectangle type.
package main

import (
    "fmt"
    "math"
)

type shape interface {
    area() float64
    perimeter() float64
}

type circle struct {
    radius float64
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}

type rectangle struct {
    length, width float64
}

func (r rectangle) area() float64 {
    return r.length * r.width
}

func (r rectangle) perimeter() float64 {
    return 2 * (r.length + r.width)
}

func main() {
    c := circle{radius: 5}
    r := rectangle{length: 5, width: 3}

    fmt.Printf("Circle - Area: %.2f, Perimeter: %.2f\n", c.area(), c.perimeter())
    fmt.Printf("Rectangle - Area: %.2f, Perimeter: %.2f\n", r.area(), r.perimeter())
}
