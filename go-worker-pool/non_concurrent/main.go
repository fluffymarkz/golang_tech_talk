package main

import (
	"fmt"
	"time"
)

func main() {
	numTasks := 15
	tasks := getTasks(numTasks)

	for i, v := range tasks {
		time.Sleep(time.Millisecond * 500)
		fmt.Println("Worker", i, "processing job `", v, "`")
		fmt.Println(fmt.Sprintf("Job `%d` done by worker `%d`", v, i))
	}

	fmt.Println("done")
}

func getTasks(nJobs int) []int {
	tasks := []int{}

	for i := 1; i <= nJobs; i++ {
		tasks = append(tasks, i)
	}

	return tasks
}
