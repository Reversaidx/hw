package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

var (
	ErrorLimit int64
	Tasks      []Task
)

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// Place your code here.
	Tasks = tasks
	ErrorLimit = int64(m)
	taskChan := make(chan Task)

	var ec int64

	wg := &sync.WaitGroup{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go consumer(taskChan, &ec, i, wg)
	}
	if err := producer(taskChan, &ec); err != nil {
		return err
	}
	wg.Wait()
	return nil
}

func producer(task chan Task, ec *int64) error {
	for _, t := range Tasks {
		if *ec >= ErrorLimit && ErrorLimit > 0 {
			close(task)
			return ErrErrorsLimitExceeded
		}
		task <- t
	}

	close(task)
	return nil
}

func consumer(task chan Task, ec *int64, i int, wg *sync.WaitGroup) {
	fmt.Printf("Thread: %v is running\n", i)
	defer wg.Done()
	for i := range task {
		err := i()
		if err != nil {
			atomic.AddInt64(ec, 1)
			fmt.Println(err)
		}
	}
	fmt.Printf("Thread: %v is stopped \n", i)
}
