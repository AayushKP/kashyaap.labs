package main

import "fmt"

func main(){
	inventory := map[string]int{
		"laptop": 5,
		"keyboard": 10,
		"mouse": 15,
	}

	fmt.Println(inventory)
	inventory["monitor"] = 7

	inventory["mouse"] = 20
	delete(inventory, "keyboard")

}
