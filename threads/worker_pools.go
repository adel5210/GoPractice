package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("W", id, "started", j)
		time.Sleep(time.Second)
		fmt.Println("W", id, "finished", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	res := make(chan int, numJobs)

	//set open workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, res)
	}

	for j := 1; j <= numJobs; j++ {
		fmt.Println("HIT")
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-res
	}
}
