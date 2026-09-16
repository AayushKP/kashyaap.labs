package main

import "fmt"

func main(){
	x:= 2
	y:= 9.322
	fmt.Printf("%T ",x)//Type
	fmt.Printf("%v ",x)//Value
	fmt.Printf("%b ",x)//Binary Value
	fmt.Printf("%e ",y)//Scientific notation
	fmt.Printf("%f ", y)//Float
	fmt.Printf("%.2f ", y)//Float
}
