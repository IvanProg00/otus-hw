package memorystorage

import (
	"context"
	"fmt"
	"time"

	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/storage"
	errorsstorage "github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/storage/errors"
	"github.com/google/uuid"
)

func (s *memoryStorage) CreateEvent(ctx context.Context, event storage.Event) error {
	s.events = append(s.events, event)
	return nil
}

func (s *memoryStorage) UpdateEvent(ctx context.Context, event storage.Event) error {
	for i, ev := range s.events {
		if ev.ID == event.ID {
			s.events[i] = storage.Event{
				ID:          ev.ID,
				Title:       event.Title,
				Description: event.Description,
				StartAt:     event.StartAt,
				FinishAt:    event.FinishAt,
				UserID:      event.UserID,
			}
			return nil
		}
	}

	return errorsstorage.ErrNotFound
}

func (s *memoryStorage) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	for i, ev := range s.events {
		if ev.ID == id {
			s.events = append(s.events[:i], s.events[i+1:]...)
			return nil
		}
	}

	return errorsstorage.ErrNotFound
}

func (s *memoryStorage) ListByDayEvent(ctx context.Context, date time.Time) ([]storage.Event, error) {
	res := []storage.Event{}

	for _, event := range s.events {
		if event.StartAt.Year() == date.Year() && event.StartAt.YearDay() == date.YearDay() {
			res = append(res, event)
		}
	}

	return res, nil
}

func (s *memoryStorage) ListByWeekEvent(ctx context.Context, startWeek time.Time) ([]storage.Event, error) {
	res := []storage.Event{}

	startWeek = time.Date(startWeek.Year(), startWeek.Month(), startWeek.Day(), 0, 0, 0, 0, startWeek.Location())
	endWeek := startWeek.AddDate(0, 0, 7)
	fmt.Println(endWeek)

	for _, event := range s.events {
		if startWeek.Before(event.StartAt) && endWeek.After(event.StartAt) {
			res = append(res, event)
		}
	}

	return res, nil
}

func (s *memoryStorage) ListByMonthEvent(ctx context.Context, date time.Time) ([]storage.Event, error) {
	res := []storage.Event{}

	startMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	endMonth := startMonth.AddDate(0, 1, 0)

	fmt.Println(date)
	fmt.Println(startMonth, endMonth)
	for _, event := range s.events {
		fmt.Println(event.StartAt)
		if startMonth.Before(event.StartAt) && endMonth.After(event.StartAt) {
			res = append(res, event)
		}
	}

	return res, nil
}
