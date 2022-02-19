package storage

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Storage interface {
	Connect(ctx context.Context) error
	Close() error

	CreateEvent(ctx context.Context, event Event) error
	UpdateEvent(ctx context.Context, event Event) error
	DeleteEvent(ctx context.Context, id uuid.UUID) error
	ListByDayEvent(ctx context.Context, date time.Time) ([]Event, error)
	ListByWeekEvent(ctx context.Context, date time.Time) ([]Event, error)
	ListByMonthEvent(ctx context.Context, date time.Time) ([]Event, error)
}
