package sqlstorage

import (
	"context"
	"fmt"

	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/config"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type Storage struct { // TODO
	hostname string
	port     int
	username string
	password string
	database string
	db       *sqlx.DB
}

func New(config config.DatabaseConf) *Storage {
	return &Storage{
		hostname: config.Hostname,
		port:     config.Port,
		username: config.Username,
		password: config.Password,
		database: config.Database,
	}
}

func (s *Storage) Connect(ctx context.Context) error {
	var err error
	s.db, err = sqlx.ConnectContext(ctx, "pgx",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
			s.hostname, s.port, s.username, s.password, s.database))
	return err
}

func (s *Storage) Close(ctx context.Context) error {
	// TODO
	return nil
}
