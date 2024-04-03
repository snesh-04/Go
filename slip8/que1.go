// Write a program in GO language to accept the book details such
// as BookID, Title, Author, Price. Read and display the details of
// ‘n’ number of books
package main

import "fmt"

type book struct {
    bookID int
    title  string
    author string
    price  float64
}

func main() {
    var n int
    fmt.Print("Enter the number of books: ")
    fmt.Scan(&n)

    books := make([]book, n)
    for i := 0; i < n; i++ {
        fmt.Printf("Enter details for book %d:\n", i+1)
        fmt.Print("Book ID: ")
        fmt.Scan(&books[i].bookID)
        fmt.Print("Title: ")
        fmt.Scan(&books[i].title)
        fmt.Print("Author: ")
        fmt.Scan(&books[i].author)
        fmt.Print("Price: ")
        fmt.Scan(&books[i].price)
    }

    fmt.Println("\nBook details:")
    for i, b := range books {
        fmt.Printf("Book %d\n", i+1)
        fmt.Printf("Book ID: %d\n", b.bookID)
        fmt.Printf("Title: %s\n", b.title)
        fmt.Printf("Author: %s\n", b.author)
        fmt.Printf("Price: %.2f\n", b.price)
        fmt.Println()
    }
}
