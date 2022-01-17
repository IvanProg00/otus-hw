package app

import (
	"context"
	"time"

	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
)

type App struct { // TODO
	Storage Storage
	Logger  Logger
}

type Logger interface { // TODO
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Debug(msg string)
}

type Storage interface { // TODO
	CreateEvent(event storage.Event) error
	UpdateEvent(id uuid.UUID, event storage.Event) error
	DeleteEvent(id uuid.UUID) error
	ListByDayEvent(date time.Time) ([]storage.Event, error)
	ListByWeekEvent(date time.Time) ([]storage.Event, error)
	ListByMonthEvent(date time.Time) ([]storage.Event, error)
}

func New(logger Logger, storage Storage) *App {
	return &App{}
}

func (a *App) CreateEvent(ctx context.Context, id uuid.UUID, title string) error {
	// TODO
	return nil
	// return a.Storage.CreateEvent(storage.Event{ID: id, Title: title})
}

// TODO
