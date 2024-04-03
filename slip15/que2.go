// Write a program in GO language to read XML file into structure
// and display structure
package main

import (
    "encoding/xml"
    "fmt"
    "os"
)

type Person struct {
    XMLName xml.Name `xml:"person"`
    Name    string   `xml:"name"`
    Age     int      `xml:"age"`
    City    string   `xml:"city"`
}

func main() {
    file, err := os.Open("person.xml")
    if err != nil {
        fmt.Println("Error opening XML file:", err)
        return
    }
    defer file.Close()

    var person Person
    decoder := xml.NewDecoder(file)
    if err := decoder.Decode(&person); err != nil {
        fmt.Println("Error decoding XML:", err)
        return
    }

    fmt.Println("XML Decoded into Structure:")
    fmt.Println("Name:", person.Name)
    fmt.Println("Age:", person.Age)
    fmt.Println("City:", person.City)
}
