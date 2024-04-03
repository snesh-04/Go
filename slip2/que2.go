// B) Write a program in GO language to print file information.

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Get the file path from user input
	var filePath string
	fmt.Println("Enter the path to the file:")
	fmt.Scanln(&filePath)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %s\n", err)
	}
	defer file.Close()

	// Get file information
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("Error getting file information: %s\n", err)
	}

	// Print file information
	fmt.Printf("File Name: %s\n", fileInfo.Name())
	fmt.Printf("Size: %d bytes\n", fileInfo.Size())
	fmt.Printf("Permissions: %s\n", fileInfo.Mode())
	fmt.Printf("Last Modified: %s\n", fileInfo.ModTime())
	fmt.Printf("Is Directory: %t\n", fileInfo.IsDir())
}
