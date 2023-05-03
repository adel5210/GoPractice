package main

import "fmt"

func main() {
	jobs := make(chan int, 3)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("receive job ", j)
			} else {
				fmt.Println("completed job")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 2; j++ {
		jobs <- j
		fmt.Println("sending job", j)
	}

	close(jobs)
	fmt.Println("sent all")

	<-done
}
