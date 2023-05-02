package main

import "fmt"

type animal interface {
	weight() int
	speciesType() int
}

type tiger struct {
}

func (t tiger) weight() int {
	return 100
}

func main() {
	t := tiger{}
	a := t.weight()
	fmt.Println(a)
}
