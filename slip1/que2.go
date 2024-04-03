// Write a program in GO language to accept n student details like roll_no,
// stud_name, mark1,mark2, mark3. Calculate the total and average of
// marks using structure.
package main

import (
	"fmt"
)

type Student struct {
	RollNo              int
	StudName            string
	Mark1, Mark2, Mark3 float64
}

func main() {
	var n int
	fmt.Println("Enter the number of students:")
	fmt.Scanln(&n)

	students := make([]Student, n)

	for i := 0; i < n; i++ {
		var rollNo int
		var studName string
		var mark1, mark2, mark3 float64

		fmt.Printf("Enter details for student %d:\n", i+1)
		fmt.Println("Roll No:")
		fmt.Scanln(&rollNo)
		fmt.Println("Student Name:")
		fmt.Scanln(&studName)
		fmt.Println("Mark 1:")
		fmt.Scanln(&mark1)
		fmt.Println("Mark 2:")
		fmt.Scanln(&mark2)
		fmt.Println("Mark 3:")
		fmt.Scanln(&mark3)

		students[i] = Student{RollNo: rollNo, StudName: studName, Mark1: mark1, Mark2: mark2, Mark3: mark3}
	}

	for _, student := range students {
		total := student.Mark1 + student.Mark2 + student.Mark3
		average := total / 3.0
		fmt.Printf("Student: %s, Roll No: %d, Total: %.2f, Average: %.2f\n", student.StudName, student.RollNo, total, average)
	}
}
