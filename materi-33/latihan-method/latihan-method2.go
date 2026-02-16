package main

import "fmt"

type Book struct {
	Title, Author string
	Pages         int
}

func (b Book) describe() {
	fmt.Printf("Book: %s by %s, %d pages\n", b.Title, b.Author, b.Pages)
}

func (b Book) isLongBook() {
	if b.Pages > 300 {
		fmt.Println("This is a long book")
	} else {
		fmt.Println("Short read")
	}
}

func main() {
	book := Book{
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Pages:  180,
	}

	book.describe()
	book.isLongBook()
}