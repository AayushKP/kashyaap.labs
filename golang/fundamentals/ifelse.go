package main

import "fmt"

func main(){
	x:=2
	if x<3{
		fmt.Println("run")
	} else if x>4{
		fmt.Println("abort")
	} else{
		fmt.Println("nothing")
	}
}
