// B) Write a program in GO language to create an interface shape that
// includes area and volume. Implements these methods in square
// and rectangle type.
package main

import (
    "fmt"
)

type shape interface {
    area() float64
    volume() float64
}

type square struct {
    side float64
}

func (s square) area() float64 {
    return s.side * s.side
}

func (s square) volume() float64 {
    return 0
}

type rectangle struct {
    length, width float64
}

func (r rectangle) area() float64 {
    return r.length * r.width
}

func (r rectangle) volume() float64 {
    return 0
}

func main() {
    s := square{side: 5}
    r := rectangle{length: 5, width: 3}

    fmt.Printf("Square - Area: %.2f, Volume: %.2f\n", s.area(), s.volume())
    fmt.Printf("Rectangle - Area: %.2f, Volume: %.2f\n", r.area(), r.volume())
}
