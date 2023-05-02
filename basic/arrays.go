package main

import "fmt"

func main() {
	var a [10]int
	a[1] = 12
	fmt.Println(a)
	fmt.Println(len(a))

	b := [5]int{1, 2, 3, 4}
	fmt.Println(b)
}
