package workerpool

import (
	"context"
	"errors"
	"runtime"

	"golang.org/x/sync/semaphore"
)

type Task[T any] func(arg T) error

type WorkerPool[T any] struct {
	maxWorkers int
	sem        *semaphore.Weighted
	ctx        context.Context
	errs       []error
}

type WorkerPoolHandler[T any] interface {
	Run(value T) error
}

func (s *WorkerPool[T]) Run(task Task[T], args ...T) error {
	s.errs = make([]error, len(args))
	for idx, arg := range args {
		if err := s.sem.Acquire(s.ctx, 1); err != nil {
			break
		}
		go func(currIdx int, currArg T) {
			defer s.sem.Release(1)

			s.errs[currIdx] = task(currArg)
		}(idx, arg)
	}

	return nil
}

func (s *WorkerPool[T]) Wait() error {
	return s.sem.Acquire(s.ctx, int64(s.maxWorkers))
}

// chack errors of runner jobs
func (s *WorkerPool[T]) Check() error {
	return errors.Join(s.errs...)
}

func NewWorkerPool[T any](numWorkers int) *WorkerPool[T] {
	numMax := runtime.GOMAXPROCS(0)
	if numWorkers > 0 {
		numMax = numWorkers
	}
	sem := semaphore.NewWeighted(int64(numMax))
	return &WorkerPool[T]{
		maxWorkers: numMax,
		sem:        sem,
		ctx:        context.TODO(),
	}
}
