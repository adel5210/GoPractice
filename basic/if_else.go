package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if num := 9; num > 0 {
		fmt.Println(num, "is less")
	}
}
