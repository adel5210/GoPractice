package main

import (
	"fmt"
	"time"
)

// a simple thread
func someMethod() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	someMethod()
	time.Sleep(time.Second)
	go someMethod()
	time.Sleep(time.Second)
}
