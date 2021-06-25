package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	results := make(chan string)
	nWorkers := 15000
	nJobs := 150000
	wg := sync.WaitGroup{}

	jobs := getJobs(nJobs)

	for i := 1; i <= nWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, wg.Done)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println(res)
	}

	// go func() {
	// 	for res := range results {
	// 		fmt.Println(res)
	// 	}
	// }()

	// wg.Wait()
	// close(results)

	fmt.Println("done")
}

func getJobs(nJobs int) <-chan int {
	jobs := make(chan int, nJobs)

	for i := 0; i < nJobs; i++ {
		jobs <- i
	}
	close(jobs)

	return jobs
}

func worker(workerID int, jobs <-chan int, results chan<- string, wgDone func()) {
	defer wgDone()
	for job := range jobs {
		// time.Sleep(time.Duration(rand.Int63n(1e3)) * time.Millisecond)
		time.Sleep(time.Millisecond * 500)
		fmt.Println("Worker", workerID, "processing job `", job, "`")
		results <- fmt.Sprintf("Job `%d` done by worker `%d`", job, workerID)
	}
}
