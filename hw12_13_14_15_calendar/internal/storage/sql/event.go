package sqlstorage

import (
	"context"
	"time"

	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/storage"
	errorsstorage "github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/storage/errors"
	"github.com/google/uuid"
)

func (s *sqlStorage) CreateEvent(ctx context.Context, event storage.Event) error {
	_, err := s.db.NamedExecContext(ctx, "INSERT INTO events (id,title,description,start_at,finish_at,user_id) VALUES"+
		" (:id,:title,:description,:start_at,:finish_at,:user_id)",
		&event)
	return err
}

func (s *sqlStorage) UpdateEvent(ctx context.Context, event storage.Event) error {
	_, err := s.db.NamedExecContext(ctx, "UPDATE events SET (title,description,start_at,finish_at,user_id) ="+
		" (:title,:description,:start_at,:finish_at,:user_id)"+
		" WHERE id=:id",
		&event)
	return err
}

func (s *sqlStorage) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	return errorsstorage.ErrNotFound
}

func (s *sqlStorage) ListByDayEvent(ctx context.Context, date time.Time) ([]storage.Event, error) {
	res := []storage.Event{}

	rows, err := s.db.Queryx("SELECT id, title FROM events")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var event storage.Event
		if err := rows.StructScan(&event); err != nil {
			return nil, err
		}
		res = append(res, event)
	}

	return res, nil
}

func (s *sqlStorage) ListByWeekEvent(ctx context.Context, startWeek time.Time) ([]storage.Event, error) {
	return []storage.Event{}, nil
}

func (s *sqlStorage) ListByMonthEvent(ctx context.Context, date time.Time) ([]storage.Event, error) {
	return []storage.Event{}, nil
}
