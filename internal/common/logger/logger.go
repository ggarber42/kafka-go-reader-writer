package logger

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

type Logger struct {
	Logger slog.Logger
}

func NewLogger() *Logger {
	w := os.Stderr
	logger := slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.DateTime,
			NoColor:    false,
		}),
	)
	slog.SetDefault(logger)
	return &Logger{*logger}
}

func (l *Logger) Info(msg string, info ...any) {
	l.Logger.Info(msg, info)
}

func (l *Logger) Error(msg string, info ...any){
	l.Logger.Error(msg, info)
}
