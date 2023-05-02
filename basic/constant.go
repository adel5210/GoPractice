package main

import "fmt"

func main() {
	const a = "no type"
	fmt.Println(a)

	const b string = "with type"
	fmt.Println(b)
}
