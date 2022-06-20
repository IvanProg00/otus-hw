package internalhttp

import (
	"context"
	"net/http"
	"time"

	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Server struct {
	Logger      Logger
	Application Application
	srv         http.Server
}

type Logger interface {
	Info(msg string)
}

type Application interface { // TODO
	CreateEvent(ctx context.Context, title, description string,
		startAt, finishAt time.Time, userID uuid.UUID) error
	UpdateEvent(ctx context.Context, id uuid.UUID, title, description string,
		startAt, finishAt time.Time, userID uuid.UUID) error
	DeleteEvent(ctx context.Context, id uuid.UUID) error
	ListByDayEvent(ctx context.Context, date time.Time) ([]storage.Event, error)
	ListByWeekEvent(ctx context.Context, date time.Time) ([]storage.Event, error)
	ListByMonthEvent(ctx context.Context, date time.Time) ([]storage.Event, error)
}

func NewServer(logger Logger, app Application) *Server {
	return &Server{
		Logger:      logger,
		Application: app,
		srv:         http.Server{},
	}
}

func (s *Server) Start(ctx context.Context, addr string) error {
	s.srv.Addr = addr
	r := mux.NewRouter()
	r.HandleFunc("/hello-world", s.helloWorld)

	http.Handle("/", s.loggingMiddleware(r))
	if err := s.srv.ListenAndServe(); err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

// TODO
