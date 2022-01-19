package storage

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID
	Title       string
	Description string
	StartAt     time.Time
	FinishAt    time.Time
	UserID      uuid.UUID
}
