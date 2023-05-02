package main

import "fmt"

func main() {
	s := make([]int, 10)
	fmt.Println(s)

	s[1] = 1
	s = append(s, 12)
	fmt.Println(s)

	c := s
	s[1] = 3
	fmt.Println(c)

	copy(c, s)
	s[1] = 2
	fmt.Println(c)
	fmt.Println(c[:3])
	fmt.Println(c[1:5])

	//remove last elem
	c = c[:len(c)-1]
	fmt.Println(c)
}
