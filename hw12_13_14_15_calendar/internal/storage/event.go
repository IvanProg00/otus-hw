package storage

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	StartAt     time.Time `db:"start_at"`
	FinishAt    time.Time `db:"finish_at"`
	UserID      uuid.UUID `db:"user_id"`
}
