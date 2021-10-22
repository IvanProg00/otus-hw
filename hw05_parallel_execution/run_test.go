package hw05parallelexecution

import (
	"errors"
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestRun(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("if were errors in first M tasks, than finished not more N+M tasks", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32

		for i := 0; i < tasksCount; i++ {
			err := fmt.Errorf("error from task %d", i)
			tasks = append(tasks, func() error {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
				atomic.AddInt32(&runTasksCount, 1)
				return err
			})
		}

		workersCount := 10
		maxErrorsCount := 23
		err := Run(tasks, workersCount, maxErrorsCount)

		require.Truef(t, errors.Is(err, ErrErrorsLimitExceeded), "actual err - %v", err)
		require.LessOrEqual(t, runTasksCount, int32(workersCount+maxErrorsCount), "extra tasks were started")
	})

	t.Run("tasks without errors", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		var sumTime time.Duration

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
			sumTime += taskSleep

			tasks = append(tasks, func() error {
				time.Sleep(taskSleep)
				atomic.AddInt32(&runTasksCount, 1)
				return nil
			})
		}

		workersCount := 5
		maxErrorsCount := 1

		start := time.Now()
		err := Run(tasks, workersCount, maxErrorsCount)
		elapsedTime := time.Since(start)
		require.NoError(t, err)

		require.Equal(t, runTasksCount, int32(tasksCount), "not all tasks were completed")
		require.LessOrEqual(t, int64(elapsedTime), int64(sumTime/2), "tasks were run sequentially?")
	})
}

func TestRun_minimumWorkerNumber(t *testing.T) {
	err := Run([]Task{}, 0, 1)
	require.ErrorIs(t, err, ErrMinWorkerNumber)
}

func TestRun_errorsLimitExceded(t *testing.T) {
	tests := []struct {
		m int
	}{
		{
			m: -1,
		}, {
			m: 0,
		}, {
			m: -483,
		},
	}

	for _, test := range tests {
		err := Run([]Task{}, 1, test.m)
		require.ErrorIs(t, err, ErrErrorsLimitExceeded)
	}
}

func TestRun_workersMoreThanTasks(t *testing.T) {
	taskCount := 10
	tasks := make([]Task, 0, taskCount)
	var sumTime time.Duration
	var taskCounter int32

	for i := 0; i < taskCount; i++ {
		taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
		sumTime += taskSleep

		tasks = append(tasks, func() error {
			time.Sleep(taskSleep)
			atomic.AddInt32(&taskCounter, 1)
			return nil
		})
	}

	start := time.Now()
	err := Run(tasks, 40, 1)
	elapsedTime := time.Since(start)

	require.NoError(t, err)
	require.Equal(t, taskCounter, int32(taskCount), "not all tasks were completed")
	require.LessOrEqual(t, int64(elapsedTime), int64(sumTime/2), "tasks were run sequentially?")
}
