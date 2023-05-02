package main

import "fmt"

func main() {
	m := make(map[string]bool)
	fmt.Println(m)

	m["a"] = true
	fmt.Println(m)
	fmt.Println(len(m))

	delete(m, "a")
	_, hasKey := m["a"]
	fmt.Println(hasKey)

	n := map[string]int{
		"test1": 1,
		"test2": 2,
	}

	fmt.Println(n)
}
