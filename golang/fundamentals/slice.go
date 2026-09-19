package main

import "fmt"

func main() {
	//slice is similar to vectors in cpp
	//Slice = a descriptor/view over an underlying array.
	numbers := []int{10, 20, 30}

	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	numbers = append(numbers,50,60)
	fmt.Println(numbers)

	//sub slice by slicing
	part:= numbers[1:4]
	fmt.Println(part)
}
