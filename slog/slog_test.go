package slog_test

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/prakashpandey/golog/log"
	"github.com/prakashpandey/golog/slog"
)

// TestSlogLogger_LogLevel tests that log levels are respected.
func TestSlogLogger_LogLevel(t *testing.T) {
	r, w := io.Pipe()
	config := log.Conf{
		Outputs:  []io.Writer{w},
		UseJSON:  false,
		LogLevel: log.Warn, // Only log warnings and errors
	}

	logger := slog.NewSlogLogger(config)

	go func() {
		defer w.Close()
		logger.Info(context.Background(), "This should not log", "key", "value")
		logger.Warn(context.Background(), "This should log", "key", "value")
	}()

	buf := new(strings.Builder)
	io.Copy(buf, r)

	if strings.Contains(buf.String(), "This should not log") {
		t.Errorf("Expected 'This should not log' to be filtered out, but it was logged")
	}
	if !strings.Contains(buf.String(), "This should log") {
		t.Errorf("Expected 'This should log' in log output, but it was not logged")
	}
}

// TestSlogLogger_DebugLevel tests the logger when log level is set to Debug.
func TestSlogLogger_DebugLevel(t *testing.T) {
	r, w := io.Pipe()
	config := log.Conf{
		Outputs:  []io.Writer{w},
		UseJSON:  false,
		LogLevel: log.Debug, // Allow all logs (Debug, Info, Warn, Error)
	}

	logger := slog.NewSlogLogger(config)

	go func() {
		defer w.Close()
		logger.Debug(context.Background(), "Debug message", "key", "value")
		logger.Info(context.Background(), "Info message", "key", "value")
		logger.Warn(context.Background(), "Warn message", "key", "value")
		logger.Error(context.Background(), "Error message", "key", "value")
	}()

	buf := new(strings.Builder)
	io.Copy(buf, r)

	if !strings.Contains(buf.String(), "Debug message") {
		t.Errorf("Expected 'Debug message' in log output, but it was not logged")
	}
	if !strings.Contains(buf.String(), "Info message") {
		t.Errorf("Expected 'Info message' in log output, but it was not logged")
	}
	if !strings.Contains(buf.String(), "Warn message") {
		t.Errorf("Expected 'Warn message' in log output, but it was not logged")
	}
	if !strings.Contains(buf.String(), "Error message") {
		t.Errorf("Expected 'Error message' in log output, but it was not logged")
	}
}
