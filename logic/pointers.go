package main

import "fmt"

func passByRef(ptr *int) {
	*ptr = 0
}

func passByVal(val int) {
	val = 0
}

func main() {
	i := 100

	passByVal(i)
	fmt.Println(i)

	passByRef(&i)
	fmt.Println(i)
}
