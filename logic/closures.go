package main

import "fmt"

// anonymous func
func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq() // new instance anonymous func
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
