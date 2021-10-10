package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

var ErrorLimit int64
var ThreadLimit int
var Tasks []Task

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// Place your code here.
	Tasks = tasks
	ErrorLimit = int64(m)
	ThreadLimit = n
	taskChan := make(chan Task)

	var ec int64

	wg := &sync.WaitGroup{}

	wg.Add(n)

	go producer(taskChan, &ec)
	for i := 0; i < n; i++ {
		go consumer(taskChan, &ec, i, wg)
	}
	wg.Wait()

	//<-done
	return nil
}
func producer(task chan Task, ec *int64) error {
	for _, t := range Tasks {
		if *ec >= ErrorLimit {

			close(task)
			//return ErrErrorsLimitExceeded
		}
		task <- t
	}

	close(task)
	return nil
}

func consumer(task chan Task, ec *int64, i int, wg *sync.WaitGroup) error {
	fmt.Printf("Thread: %v is running\n", i)
	defer wg.Done()
	for i := range task {
		err := i()
		if err != nil {
			atomic.AddInt64(ec, 1)
			fmt.Println(err)
		}
	}
	fmt.Printf("Thread: %v is Stopped without error\n", i)

	return nil
}
