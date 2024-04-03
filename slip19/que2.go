// Write a program in the GO language program to open a file in
// READ only mode.
package main

import (
    "fmt"
    "os"
)

func main() {
    fileName := "example.txt"
    file, err := os.Open(fileName)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    fmt.Printf("Successfully opened file %s\n", fileName)
    // You can now read from the file using 'file' variable
}
