package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 1
	f := float64(i)
	a := "1234hello"
	//strconv used to manipulate strings
	x, err := strconv.Atoi(a)
	fmt.Println(x,err)



	fmt.Printf("Type %T value %v", f, f)
}
