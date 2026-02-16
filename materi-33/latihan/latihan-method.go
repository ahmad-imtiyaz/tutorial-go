package main

import "fmt"

type Book struct {
	Title, Author string
	Pages int
}

func main () {
	Book := Book{
		Title: "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Pages: 180,
	}
	fmt.Println("Book:", Book.Title,"by", Book.Author, Book.Pages,"pages.")
	if Book.Pages > 300 {
		fmt.Println("This book is a long read.")
	} else {
		fmt.Println("Short read.")
	}

}