package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

// display is a METHOD of the Book struct.
// (b Book) is called the RECEIVER.
// It means this method belongs to the Book type.
// b represents the particular Book on which we call the method.
func (b Book) display() {
	fmt.Println("Title:", b.Title)
	fmt.Println("Author:", b.Author)
	fmt.Println("Pages:", b.Pages)
}

// The receiver (b Book) is what makes display()
// a method associated with the Book type.

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

	// Call the method on book1
	book1.display()
}
