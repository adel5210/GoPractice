package main

import "fmt"

func main() {
	messages := make(chan string)
	go func() { messages <- "data" }()
	msg := <-messages
	fmt.Println(msg)
	fmt.Println(msg)
}
