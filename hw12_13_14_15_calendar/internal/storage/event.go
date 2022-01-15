package storage

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID             uuid.UUID
	Title          string
	Description    string
	DateTime       time.Time
	FinishDateTime time.Time
	UserID         uuid.UUID
}
