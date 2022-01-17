package sqlstorage

import (
	"time"

	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/storage"
	errorsstorage "github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/storage/errors"
	"github.com/google/uuid"
)

func (s *Storage) CreateEvent(event storage.Event) error {
	return nil
}

func (s *Storage) UpdateEvent(id uuid.UUID, event storage.Event) error {
	return errorsstorage.ErrNotFound
}

func (s *Storage) DeleteEvent(id uuid.UUID) error {
	return errorsstorage.ErrNotFound
}

func (s *Storage) ListByDayEvent(date time.Time) ([]storage.Event, error) {
	return []storage.Event{}, nil
}

func (s *Storage) ListByWeekEvent(startWeek time.Time) ([]storage.Event, error) {
	return []storage.Event{}, nil
}

func (s *Storage) ListByMonthEvent(date time.Time) ([]storage.Event, error) {
	return []storage.Event{}, nil
}
