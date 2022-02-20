package logger

import (
	"fmt"
	"time"
)

type loggerLevel int8

const (
	INFO loggerLevel = iota
	WARN
	ERROR
	DEBUG
)

type logger struct {
	level loggerLevel
}

type Logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Debug(msg string)
}

func New(level string) Logger {
	logger := logger{}

	switch level {
	case "info":
		logger.level = INFO
	case "warn":
		logger.level = WARN
	case "error":
		logger.level = ERROR
	default:
		logger.level = DEBUG
	}

	return &logger
}

func (l logger) Info(msg string) {
	if l.level >= INFO {
		log(msg, "INFO")
	}
}

func (l logger) Warn(msg string) {
	if l.level >= WARN {
		log(msg, "WARN")
	}
}

func (l logger) Error(msg string) {
	if l.level >= ERROR {
		log(msg, "ERROR")
	}
}

func (l logger) Debug(msg string) {
	if l.level >= DEBUG {
		log(msg, "DEBUG")
	}
}

func log(msg, level string) {
	fmt.Printf("%s [%s] - %s\n", time.Now().Format("2006-01-02 15:04:05"), level, msg)
}
