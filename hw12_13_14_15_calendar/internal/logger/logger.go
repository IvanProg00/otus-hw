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

type Logger struct {
	level loggerLevel
}

func New(level string) *Logger {
	logger := Logger{}

	switch level {
	case "warn":
		logger.level = WARN
	case "error":
		logger.level = ERROR
	case "debug":
		logger.level = DEBUG
	default:
		logger.level = INFO
	}

	return &logger
}

func (l Logger) Info(msg string) {
	if l.level >= INFO {
		log(msg, "INFO")
	}
}

func (l Logger) Warn(msg string) {
	if l.level >= WARN {
		log(msg, "WARN")
	}
}

func (l Logger) Error(msg string) {
	if l.level >= ERROR {
		log(msg, "ERROR")
	}
}
func (l Logger) Debug(msg string) {
	if l.level >= DEBUG {
		log(msg, "DEBUG")
	}
}

func log(msg, level string) {
	fmt.Printf("%s [%s] - %s", time.Now().Format("2006-01-02 15:04:05"), level, msg)
}
