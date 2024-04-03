// Write a program in GO language to create structure student. Writea
// method show() whose receiver is a pointer of struct student. 
package main

import "fmt"

type student struct {
    name  string
    age   int
    grade string
}

func (s *student) show() {
    fmt.Printf("Name: %s\nAge: %d\nGrade: %s\n", s.name, s.age, s.grade)
}

func main() {
    s := &student{name: "Alice", age: 20, grade: "A"}
    s.show()
}
