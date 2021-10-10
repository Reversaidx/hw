package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// Place your code here.
	//var err error
	errorLimit := int64(m)
	taskChan := make(chan Task)
	errChan := make(chan error)
	doneChan := make(chan bool)
	var ec int64

	wg := &sync.WaitGroup{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go consumer(taskChan, &ec, i, wg)
	}
	go producer(taskChan, &ec, tasks, errorLimit, errChan, doneChan)
	switch {

	}

	wg.Wait()

	select {
	case <-doneChan:
		return nil
	case err := <-errChan:
		return err
	}

}

func producer(task chan Task, ec *int64, tasks []Task, errorLimit int64, errChan chan error, doneChan chan bool) {
	for _, t := range tasks {
		if atomic.LoadInt64(ec) >= errorLimit && errorLimit > 0 {
			close(task)
			errChan <- ErrErrorsLimitExceeded
			return
		}
		task <- t
	}

	close(task)
	doneChan <- true

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
