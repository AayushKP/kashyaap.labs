package main

import "fmt"

// Speaker requires a Speak() method.
type Speaker interface {
	Speak()
}

type Dog struct {
	Name string
}

// Pointer receiver.
//
// Speak() belongs to *Dog.
func (d *Dog) Speak() {
	fmt.Println(d.Name, "says: Woof!")
}

func main() {

	dog := Dog{
		Name: "Rocky",
	}

	// This works because &dog is a *Dog.
	var speaker Speaker = &dog

	speaker.Speak()

	// This would NOT compile:
	//
	// var speaker Speaker = dog
	//
	// Because Dog itself does not have Speak()
	// in its method set when Speak has a pointer receiver.
}
