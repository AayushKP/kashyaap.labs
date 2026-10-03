package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		err := errors.New("cannot divide by 0")
		return 0, err
	}
	result := a / b
	return result, nil
}

func main() {
	result,err := divide(10, 0)
	fmt.Println(result)
	fmt.Println(err)
}
