package main

import "fmt"

type base struct {
	name string
}

func (b base) someBaseMethod() {
	fmt.Println("Im base method")
}

type derived struct {
	base
	email string
}

func main() {
	d := derived{base: base{name: "test"}, email: "some_email"}
	fmt.Println(d)
}
