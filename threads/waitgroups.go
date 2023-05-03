package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Println("W", id, "starting...")
	time.Sleep(time.Second)
	fmt.Println("W", id, "done")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait()
}
