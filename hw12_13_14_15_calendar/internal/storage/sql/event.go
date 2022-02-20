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
	res, err := s.db.NamedExecContext(ctx, "UPDATE events SET (title,description,start_at,finish_at,user_id) ="+
		" (:title,:description,:start_at,:finish_at,:user_id)"+
		" WHERE id=:id",
		&event)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errorsstorage.ErrNotFound
	}

	return err
}

func (s *sqlStorage) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	res, err := s.db.ExecContext(ctx, "DELETE FROM events"+
		" WHERE id=$1",
		id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errorsstorage.ErrNotFound
	}

	return err
}

func (s *sqlStorage) ListByDayEvent(ctx context.Context, date time.Time) ([]storage.Event, error) {
	res := []storage.Event{}
	date = date.Truncate(24 * time.Hour)

	if err := s.db.SelectContext(ctx, &res, "SELECT * FROM events"+
		" WHERE date_trunc('day',start_at)=$1", date); err != nil {
		return res, err
	}

	return res, nil
}

func (s *sqlStorage) ListByWeekEvent(ctx context.Context, date time.Time) ([]storage.Event, error) {
	res := []storage.Event{}
	startWeek := date.Truncate(24 * time.Hour)
	finishWeek := startWeek.Local().AddDate(0, 0, 6)

	err := s.db.SelectContext(ctx, &res, "SELECT * FROM events"+
		" WHERE date_trunc('day',start_at)>=$1 AND date_trunc('day',start_at)<$2", startWeek, finishWeek)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *sqlStorage) ListByMonthEvent(ctx context.Context, date time.Time) ([]storage.Event, error) {
	res := []storage.Event{}
	startMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	finishMonth := startMonth.Local().AddDate(0, 1, 0)

	err := s.db.SelectContext(ctx, &res, "SELECT * FROM events"+
		" WHERE date_trunc('day',start_at)>=$1 AND date_trunc('day',start_at)<$2", startMonth, finishMonth)
	if err != nil {
		return nil, err
	}

	return res, nil
}
