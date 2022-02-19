package memorystorage

import (
	"context"
	"sync"

	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/storage"
)

type memoryStorage struct {
	events []storage.Event
	mu     sync.RWMutex
}

func New() storage.Storage {
	return &memoryStorage{
		events: []storage.Event{},
		mu:     sync.RWMutex{},
	}
}

func (s *memoryStorage) Connect(ctx context.Context) error {
	return nil
}

func (s *memoryStorage) Close() error {
	return nil
}
