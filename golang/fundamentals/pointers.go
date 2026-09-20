package main

import "fmt"

func main() {
	// x is a normal integer variable.
	x := 10

	// &x means "get the memory address of x".
	//
	// p is therefore a pointer to x.
	p := &x

	// p contains the address of x.
	fmt.Println("Address stored in p:", p)

	// *p means "go to the address stored in p
	// and get the value stored there".
	fmt.Println("Value of x through pointer:", *p)

	// We can modify x through the pointer.
	//
	// *p = 50 means:
	// "Go to the memory location p points to
	// and change its value to 50."
	*p = 50

	// x has now changed because p points to x.
	fmt.Println("x:", x)

	// *p also gives us the same updated value.
	fmt.Println("Value through pointer:", *p)
}
