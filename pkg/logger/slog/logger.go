package slog

import (
	"fmt"
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func (l *Logger) Errorf(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	l.Error(message)
}

func (l *Logger) Infof(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	l.Info(message)
}

func (l *Logger) Warnf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	l.Warn(message)
}

func NewLogger() *Logger {
	return &Logger{
		slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})),
	}
}
