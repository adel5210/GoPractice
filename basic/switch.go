package main

import "fmt"

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
		fmt.Println("case 2b")
	case 3:
		fmt.Println("case 3")
	default:
		fmt.Println("defaullt")
	}

	someVar := func(i interface{}) {
		switch t:=i.(type) {
		case bool:
			fmt.Println(t, " a bool")
		case int:
			fmt.Println(t, " an int")
		}
	}

	someVar(true)
}
