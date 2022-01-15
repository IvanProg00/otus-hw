package storage

import (
	"time"

	"github.com/google/uuid"
)

type Notice struct {
	ID       uuid.UUID
	Title    string
	DateTime time.Time
	UserID   uuid.UUID
}
