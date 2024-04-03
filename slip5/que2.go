// Write a program in GO language to accept n records of employee
// information (eno,ename,salary) and display records of employees
// having minimum salary.
package main

import (
    "fmt"
)

type Employee struct {
    eno    int
    ename  string
    salary float64
}

func main() {
    var n int
    fmt.Print("Enter the number of employees: ")
    fmt.Scan(&n)

    employees := make([]Employee, n)
    minSalary := 1e9 // Initialize to a very high value

    for i := 0; i < n; i++ {
        fmt.Printf("Enter details for employee %d:\n", i+1)
        fmt.Print("Employee Number: ")
        fmt.Scan(&employees[i].eno)
        fmt.Print("Employee Name: ")
        fmt.Scan(&employees[i].ename)
        fmt.Print("Employee Salary: ")
        fmt.Scan(&employees[i].salary)

        if employees[i].salary < minSalary {
            minSalary = employees[i].salary
        }
    }

    fmt.Printf("\nEmployees with minimum salary (%.2f):\n", minSalary)
    for _, emp := range employees {
        if emp.salary == minSalary {
            fmt.Printf("Employee Number: %d, Employee Name: %s, Salary: %.2f\n", emp.eno, emp.ename, emp.salary)
        }
    }
}
