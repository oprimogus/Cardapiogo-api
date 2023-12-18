package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

// Logger Struct
type Logger struct {
	entry *logrus.Entry
}

// NewLogger return a pointer of logger
func NewLogger(p string) *Logger {
	log := logrus.New()
	log.Out = os.Stdout

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	entry := log.WithFields(logrus.Fields{
		"service": p,
	})

	return &Logger{
		entry: entry,
	}
}

// GetLogger return a pointer of logger
func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}

// Create Non-Formatted Logs

// Debug : Create Non-Formatted Logs for debug
func (l *Logger) Debug(v ...interface{}) {
	l.entry.Debug(v...)
}

// Info : Create Non-Formatted Logs for info
func (l *Logger) Info(v ...interface{}) {
	l.entry.Info(v...)
}

// Warning : Create Non-Formatted Logs for warning
func (l *Logger) Warning(v ...interface{}) {
	l.entry.Warn(v...)
}

// Error : Create Non-Formatted Logs for error
func (l *Logger) Error(v ...interface{}) {
	l.entry.Error(v...)
}

// Métodos para criar logs formatados

// Debugf : Create Formatted Logs for debug
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.entry.Debugf(format, v...)
}

// Infof : Create Formatted Logs for info
func (l *Logger) Infof(format string, v ...interface{}) {
	l.entry.Infof(format, v...)
}

// Warningf : Create Formatted Logs for warning
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.entry.Warnf(format, v...)
}

// Errorf : Create Formatted Logs for error
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.entry.Errorf(format, v...)
}
