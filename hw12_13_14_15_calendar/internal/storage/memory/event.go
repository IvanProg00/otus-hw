package memorystorage

import (
	"time"

	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/storage"
	errorsstorage "github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/storage/errors"
	"github.com/google/uuid"
)

func (s *Storage) CreateEvent(event storage.Event) error {
	s.events = append(s.events, event)
	return nil
}

func (s *Storage) UpdateEvent(id uuid.UUID, event storage.Event) error {
	for i, ev := range s.events {
		if ev.ID == id {
			s.events[i] = event
			return nil
		}
	}

	return errorsstorage.ErrNotFound
}

func (s *Storage) DeleteEvent(id uuid.UUID) error {
	for i, ev := range s.events {
		if ev.ID == id {
			s.events = append(s.events[:i], s.events[i+1:]...)
			return nil
		}
	}

	return errorsstorage.ErrNotFound
}

func (s *Storage) ListForDayEvent(date time.Time) ([]storage.Event, error) {
	return []storage.Event{}, nil
}

func (s *Storage) ListForWeekEvent(date time.Time) ([]storage.Event, error) {
	return s.events, nil
}

func (s *Storage) ListForMonthEvent(date time.Time) ([]storage.Event, error) {
	return s.events, nil
}
