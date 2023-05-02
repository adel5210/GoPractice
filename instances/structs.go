package main

import "fmt"

type person struct {
	name  string `default:"UNKNOWN"`
	email string
}

func newPerson(name string) *person {
	return &person{name: name}
}

func (p *person) updateName(name string) {
	p.name = name
}

func main() {
	me := person{name: "test", email: "some_email"}
	fmt.Println(me)
	fmt.Println(newPerson("test2"))

	me.updateName("test3")
	fmt.Println(me)
}
