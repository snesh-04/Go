// Write a program in GO language to add or append content at the
// end of a text file
package main

import (
    "fmt"
    "os"
)

func appendToFile(fileName, content string) error {
    file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.WriteString(content)
    if err != nil {
        return err
    }

    return nil
}

func main() {
    fileName := "example.txt"
    content := "This is new content added to the file."

    err := appendToFile(fileName, content)
    if err != nil {
        fmt.Println("Error appending to file:", err)
        return
    }

    fmt.Println("Content appended to file successfully.")
}
