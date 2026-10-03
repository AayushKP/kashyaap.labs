package main

import "fmt"

// Speaker is an interface.
//
// It says:
// "Any type that has a Speak() method can be treated as a Speaker."
type Speaker interface {
	Speak()
}

// Dog is one concrete type.
type Dog struct {
	Name string
}

// Dog has its own implementation of Speak().
//
// Because Dog has Speak(), it automatically satisfies Speaker.
func (d Dog) Speak() {
	fmt.Println(d.Name, "says: Woof!")
}

// Cat is another completely different type.
type Cat struct {
	Name string
}

// Cat also has a Speak() method.
//
// Therefore Cat also automatically satisfies Speaker.
func (c Cat) Speak() {
	fmt.Println(c.Name, "says: Meow!")
}

// This function does NOT care whether it receives a Dog or Cat.
//
// It only cares that the value satisfies Speaker,
// meaning it must have a Speak() method.
func makeSound(s Speaker) {
	s.Speak()
}

func main() {

	// Create a Dog object.
	dog := Dog{
		Name: "Rocky",
	}

	// Create a Cat object.
	cat := Cat{
		Name: "Milo",
	}

	// The same function can work with Dog.
	makeSound(dog)

	// The same function can also work with Cat.
	makeSound(cat)
}
