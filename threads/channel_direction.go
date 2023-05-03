package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pongs chan<- string, pings <-chan string) {
	tmp := <-pings
	pongs <- tmp
}

func main() {
	p1 := make(chan string, 1)
	p2 := make(chan string, 1)

	ping(p1, "data")
	pong(p2, p1)
	fmt.Println(<-p2)
}
