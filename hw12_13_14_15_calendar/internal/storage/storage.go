package storage

import (
	"time"

	"github.com/google/uuid"
)

type StorageApi interface {
	CreateEvent(event Event) error
	UpdateEvent(id uuid.UUID, event Event) error
	DeleteEvent(id uuid.UUID) error
	ListForDayEvent(date time.Time) ([]Event, error)
	ListForWeekEvent(date time.Time) ([]Event, error)
	ListForMonthEvent(date time.Time) ([]Event, error)
}
