package storage

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID             uuid.UUID
	Title          string `faker:"len=50"`
	Description    string
	DateTime       time.Time
	FinishDateTime time.Time
	UserID         uuid.UUID
}
