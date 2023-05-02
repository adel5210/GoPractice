package main

import "fmt"

func main() {
	someArr := []int{1, 2, 3, 4}

	for index, val := range someArr {
		fmt.Println(index, val)
	}
}
