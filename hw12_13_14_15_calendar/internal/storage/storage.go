package storage

import (
	"time"

	"github.com/google/uuid"
)

type StorageApi interface {
	CreateEvent(event Event) error
	UpdateEvent(id uuid.UUID, event Event) error
	DeleteEvent(id uuid.UUID) error
	ListByDayEvent(date time.Time) ([]Event, error)
	ListByWeekEvent(date time.Time) ([]Event, error)
	ListByMonthEvent(date time.Time) ([]Event, error)
}
