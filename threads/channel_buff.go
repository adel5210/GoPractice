package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "data1"
	messages <- "data2"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
