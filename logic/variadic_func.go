package main

import "fmt"

func disp(somArr ...int) {
	fmt.Println(somArr, " ")
}

func main() {
	disp(1, 2)
	nums := []int{2, 4, 6, 8}
	disp(nums...)
}
