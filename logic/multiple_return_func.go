package main

import "fmt"

func vals() (string, bool) {
	return "It works", true
}

func main() {
	str, check := vals()
	fmt.Println(str)
	fmt.Println(check)
}
