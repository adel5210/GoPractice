package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.NewTimer(3000)
	<-t1.C
	fmt.Println("Timer 1 fired")

	t2 := time.NewTimer(2000)
	go func() {
		<-t2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := t2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2000)
}
