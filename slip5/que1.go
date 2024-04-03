// Write a program in GO language program to create Text file
package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    // Content to write to the file
    content := []byte("Hello, this is a text file created using Go!")

    // Write content to a file named "example.txt"
    err := ioutil.WriteFile("example.txt", content, 0644)
    if err != nil {
        fmt.Println("Error writing file:", err)
        return
    }

    fmt.Println("Text file created successfully!")
}
