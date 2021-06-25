package main

import (
	"fmt"
	"time"
)

func main() {
	results := make(chan string, 15)
	jobs := getJobs()
	nWorkers := 3

	for i := 1; i <= nWorkers; i++ {
		go worker(i, jobs, results)
	}

	for i := 0; i < 15; i++ {
		fmt.Println(<-results)
	}

	fmt.Println("done")
}

func getJobs() <-chan int {
	jobs := make(chan int, 15)

	for i := 0; i < 15; i++ {
		jobs <- i
	}
	close(jobs)

	return jobs
}

func worker(workerID int, jobs <-chan int, results chan<- string) {
	for job := range jobs {
		// time.Sleep(time.Duration(rand.Int63n(1e3)) * time.Millisecond)
		time.Sleep(time.Millisecond * 500)
		fmt.Println("Worker", workerID, "processing job `", job, "`")
		results <- fmt.Sprintf("Job `%d` done by worker `%d`", job, workerID)
	}
}
