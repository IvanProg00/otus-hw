package sqlstorage

import (
	"context"
	"fmt"

	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/jmoiron/sqlx"
)

type sqlStorage struct {
	hostname string
	port     int
	username string
	password string
	database string
	db       *sqlx.DB
}

func New(config config.DatabaseConf) storage.Storage {
	return &sqlStorage{
		hostname: config.Hostname,
		port:     config.Port,
		username: config.Username,
		password: config.Password,
		database: config.Database,
	}
}

func (s *sqlStorage) Connect(ctx context.Context) error {
	var err error
	s.db, err = sqlx.ConnectContext(ctx, "pgx",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
			s.hostname, s.port, s.username, s.password, s.database))
	return err
}

func (s *sqlStorage) Close() error {
	return s.db.Close()
}
