package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2000)
		c1 <- "r1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1000):
		fmt.Println("too late")
	}
}
