// Write a Program in GO language to accept n records of employee
// information (eno,ename,salary) and display record of employees
// having maximum salary.
package main

import (
	"fmt"
)

type Employee struct {
	Eno    int
	Ename  string
	Salary float64
}

func main() {
	var n int
	fmt.Println("Enter the number of employees:")
	fmt.Scanln(&n)

	employees := make([]Employee, n)

	// Accept employee information
	for i := 0; i < n; i++ {
		var eno int
		var ename string
		var salary float64

		fmt.Printf("Enter details for employee %d:\n", i+1)
		fmt.Println("Employee Number:")
		fmt.Scanln(&eno)
		fmt.Println("Employee Name:")
		fmt.Scanln(&ename)
		fmt.Println("Salary:")
		fmt.Scanln(&salary)

		employees[i] = Employee{Eno: eno, Ename: ename, Salary: salary}
	}

	// Find maximum salary
	maxSalary := employees[0].Salary
	for _, emp := range employees {
		if emp.Salary > maxSalary {
			maxSalary = emp.Salary
		}
	}

	// Display employees with maximum salary
	fmt.Printf("Employees with maximum salary (%.2f):\n", maxSalary)
	for _, emp := range employees {
		if emp.Salary == maxSalary {
			fmt.Printf("Employee Number: %d, Employee Name: %s, Salary: %.2f\n", emp.Eno, emp.Ename, emp.Salary)
		}
	}
}
