package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

type Logger struct {
	entry *slog.Logger
}

func NewLogger(p string) *Logger {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	log = log.With(slog.String("service", p))
	return &Logger{
		entry: log,
	}
}

func (l *Logger) GetEntry() *slog.Logger {
	return l.entry
}

func (l *Logger) WithContext(ctx context.Context) {
	log := l.entry

	if transactionId, ok := ctx.Value("transactionId").(string); ok {
		log.With(slog.String("transactionId", transactionId))
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.entry.Debug(fmt.Sprint(v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.entry.Info(fmt.Sprint(v...))
}

func (l *Logger) Warning(v ...interface{}) {
	l.entry.Warn(fmt.Sprint(v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.entry.Error(fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.entry.Debug(fmt.Sprintf(format, v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.entry.Info(fmt.Sprintf(format, v...))
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.entry.Warn(fmt.Sprintf(format, v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.entry.Error(fmt.Sprintf(format, v...))
}
