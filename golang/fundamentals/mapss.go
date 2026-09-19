package main

import "fmt"

func main() {
	//map[keytype]valuetype
	// var ages map[string]int
	//Above creates a nil map
	//keys are string values are int

	//This creates initialised map
	marks := make(map[string]int)
	marks["Aayush"] = 22
	marks["Rahul"] = 24

	fmt.Println(marks["Aayush"])

	//make() initializes certain Go built-in data structures so they are ready to be used.
	//m := make(map[string]int)
	//s := make([]int, 5)

	//We can create and initialise too
	numbers:= map[string]int{
		"Aayu":22,
		"sam": 24,
	}

	fmt.Println(numbers)

	//To check if something exists
	mark,exists:= marks["Someone"]
	//coz default value is 0 for not found keys
	if exists{
		fmt.Print(mark)
	}

	//Looping through a map
	for key,value:= range marks{
		fmt.Println(key,value)
	}

	delete(marks,"Aayush")
}
