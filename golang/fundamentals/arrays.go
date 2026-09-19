package main

import "fmt"

func main(){
	//declaring an array
	var num [5]int
	fmt.Println(num)

	//intialisation and decalration
	numbers:=[5]int{10,20,30,40,50}
	fmt.Println(numbers)
	for _,n:=range numbers{
		fmt.Printf("%d ",n)
	}
}
