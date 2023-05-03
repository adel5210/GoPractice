package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500000)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("done")
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600000)
	ticker.Stop()

	done <- true
	fmt.Println("Ticker stopped")
}
