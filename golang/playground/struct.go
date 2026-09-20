package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func main() {
	book1 := Book{
		Title:  "The Go Programming Language",
		Author: "Hirex",
		Pages:  3,
	}

	book2 := Book{
		Title:  "C",
		Author: "Abc",
		Pages:  10,
	}

	// Access a field
	fmt.Println(book1.Title)

	// Modify a field
	book1.Pages = 5
	fmt.Println("Updated pages:", book1.Pages)

	// Slice of structs
	books := []Book{book1, book2}

	// Loop through the books
	for _, book := range books {
		fmt.Println(book.Title)
	}
}
