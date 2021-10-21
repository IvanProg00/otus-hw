package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrWorkerMinNumber     = errors.New("minimum must be 1 worker")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n < 1 {
		return ErrWorkerMinNumber
	}

	wg := sync.WaitGroup{}
	ch := make(chan Task, n)
	var errorCounter int64 = int64(m)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(ch, &wg, &errorCounter)
	}

	for _, t := range tasks {
		ch <- t
	}

	close(ch)
	wg.Wait()

	if errorCounter <= 0 {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func worker(ch <-chan Task, wg *sync.WaitGroup, errorCounter *int64) {
	defer wg.Done()
	for t := range ch {
		if atomic.LoadInt64(errorCounter) > 0 {
			if err := t(); err != nil {
				atomic.AddInt64(errorCounter, -1)
			}
		}
	}
}
