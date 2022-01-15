package memorystorage

import (
	"sync"

	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	events []storage.Event
	mu     sync.RWMutex
}

func New() storage.StorageApi {
	return &Storage{
		events: []storage.Event{},
	}
}
