package memorystorage

import (
	"fmt"
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

func (s *Storage) ListByDayEvent(date time.Time) ([]storage.Event, error) {
	res := []storage.Event{}

	for _, event := range s.events {
		if event.DateTime.Year() == date.Year() && event.DateTime.YearDay() == date.YearDay() {
			res = append(res, event)
		}
	}

	return res, nil
}

func (s *Storage) ListByWeekEvent(startWeek time.Time) ([]storage.Event, error) {
	res := []storage.Event{}

	startWeek = time.Date(startWeek.Year(), startWeek.Month(), startWeek.Day(), 0, 0, 0, 0, startWeek.Location())
	endWeek := startWeek.AddDate(0, 0, 7)
	fmt.Println(endWeek)

	for _, event := range s.events {
		if startWeek.Before(event.DateTime) && endWeek.After(event.DateTime) {
			res = append(res, event)
		}
	}

	return res, nil
}

func (s *Storage) ListByMonthEvent(date time.Time) ([]storage.Event, error) {
	return s.events, nil
}
