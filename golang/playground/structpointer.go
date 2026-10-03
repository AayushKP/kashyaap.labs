package main

import "fmt"

type User struct{
	Name string
	Age int
}

func main(){
	//Create a normal User Value
	user:= User{
		Name: "Aayush",
		Age: 22,
	}

	//Get the address of user
	// p is now a pointer to the User struct
	p := &user

	// Access fields through the pointer
	// Go automatically dereferences the pointer for struct fields
	fmt.Println(p.Name)
	fmt.Println(p.Age)

	//Modify the original user through the pointer
	p.Age= 23
	fmt.Println(user.Age)
}
