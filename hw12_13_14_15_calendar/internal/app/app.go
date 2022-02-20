package app

import (
	"context"
	"time"

	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
)

type App struct {
	Storage storage.Storage
	Logger  logger.Logger
}

func New(logger logger.Logger, storage storage.Storage) *App {
	return &App{
		Logger:  logger,
		Storage: storage,
	}
}

func (a *App) CreateEvent(ctx context.Context, title, description string, startAt, finishAt time.Time,
	userID uuid.UUID) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	return a.Storage.CreateEvent(ctx, storage.Event{
		ID:          id,
		Title:       title,
		Description: description,
		StartAt:     startAt,
		FinishAt:    finishAt,
		UserID:      userID,
	})
}

func (a *App) UpdateEvent(ctx context.Context, id uuid.UUID, title, description string,
	startAt, finishAt time.Time, userID uuid.UUID) error {
	return a.Storage.UpdateEvent(ctx, storage.Event{
		ID:          id,
		Title:       title,
		Description: description,
		StartAt:     startAt,
		FinishAt:    finishAt,
		UserID:      userID,
	})
}

func (a *App) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	return a.Storage.DeleteEvent(ctx, id)
}

func (a *App) ListByDayEvent(ctx context.Context, date time.Time) ([]storage.Event, error) {
	return a.Storage.ListByDayEvent(ctx, date)
}

func (a *App) ListByWeekEvent(ctx context.Context, date time.Time) ([]storage.Event, error) {
	return a.Storage.ListByWeekEvent(ctx, date)
}

func (a *App) ListByMonthEvent(ctx context.Context, date time.Time) ([]storage.Event, error) {
	return a.Storage.ListByMonthEvent(ctx, date)
}
