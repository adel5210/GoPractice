package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg < 0 {
		return 0, errors.New("It cannot be less than zero")
	}
	return arg, nil
}

func main() {
	fmt.Println(f1(10))
	fmt.Println(f1(-10))
}
