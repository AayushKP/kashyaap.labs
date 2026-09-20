package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

// VALUE RECEIVER
//
// b is a copy of the Book.
// Changes made to b do NOT modify the original Book.
func (b Book) changeTitleValue() {
	b.Title = "Changed using value receiver"
}

// POINTER RECEIVER
//
// b is a pointer to the original Book.
// Changes made through b modify the original Book.
func (b *Book) changeTitlePointer() {
	b.Title = "Changed using pointer receiver"
}

func main() {

	book := Book{
		Title:  "The Go Programming Language",
		Author: "Hirex",
		Pages:  500,
	}

	fmt.Println("Original:", book.Title)

	// Call value receiver.
	book.changeTitleValue()

	// Original is NOT changed.
	fmt.Println("After value receiver:", book.Title)

	// Call pointer receiver.
	book.changeTitlePointer()

	// Original IS changed.
	fmt.Println("After pointer receiver:", book.Title)
}
