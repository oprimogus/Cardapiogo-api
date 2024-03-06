package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

// Logger Struct
type Logger struct {
	entry *slog.Logger
}

// NewLogger return a pointer of logger
func NewLogger(p string, ctx context.Context) *Logger {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	transactionId := ctx.Value("transactionId").(string)

	log = log.With(
		slog.String("service", p),
		slog.String("transactionId", transactionId),
	)

	return &Logger{
		entry: log,
	}
}

// NewLoggerDefault return a pointer of logger without context
func NewLoggerDefault(p string) *Logger {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	log = log.With(
		slog.String("service", p),
	)

	return &Logger{
		entry: log,
	}
}

// GetLogger return a pointer of logger
func GetLogger(p string, ctx context.Context) *Logger {
	logger := NewLogger(p, ctx)
	return logger
}

// GetLoggerDefault return a pointer of logger without context
func GetLoggerDefault(p string) *Logger {
	logger := NewLoggerDefault(p)
	return logger
}

// GetEntry retorna *logrus.Entry
func (l *Logger) GetEntry() *slog.Logger {
	return l.entry
}

// Create Non-Formatted Logs

// Debug : Create Non-Formatted Logs for debug
func (l *Logger) Debug(v ...interface{}) {
	l.entry.Debug(fmt.Sprint(v...))
}

// Info : Create Non-Formatted Logs for info
func (l *Logger) Info(v ...interface{}) {
	l.entry.Info(fmt.Sprint(v...))
}

// Warning : Create Non-Formatted Logs for warning
func (l *Logger) Warning(v ...interface{}) {
	l.entry.Warn(fmt.Sprint(v...))
}

// Error : Create Non-Formatted Logs for error
func (l *Logger) Error(v ...interface{}) {
	l.entry.Error(fmt.Sprint(v...))
}

// Métodos para criar logs formatados

// Debugf : Create Formatted Logs for debug
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.entry.Debug(fmt.Sprintf(format, v...))
}

// Infof : Create Formatted Logs for info
func (l *Logger) Infof(format string, v ...interface{}) {
	l.entry.Info(fmt.Sprintf(format, v...))
}

// Warningf : Create Formatted Logs for warning
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.entry.Warn(fmt.Sprintf(format, v...))
}

// Errorf : Create Formatted Logs for error
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.entry.Error(fmt.Sprintf(format, v...))
}
