package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func main() {
	resP := add(1, 2)
	resM := minus(2, 1)

	fmt.Println(resP)
	fmt.Println(resM)
}
