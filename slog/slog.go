package slog

import (
	"context"
	"io"
	"log/slog"
	"os"

	"github.com/prakashpandey/golog/log"
)

// SlogLogger is a concrete implementation of the Logger interface using slog.
type SlogLogger struct {
	logger   *slog.Logger
	logLevel log.Level
}

// Convert the custom LogLevel to slog's LogLevel.
func convertLogLevel(level log.Level) slog.Level {
	switch level {
	case log.Debug:
		return slog.LevelDebug
	case log.Info:
		return slog.LevelInfo
	case log.Warn:
		return slog.LevelWarn
	case log.Error:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// NewSlogLogger initializes the SlogLogger with the given config.
func NewSlogLogger(config log.Config) log.Logger {
	multiWriter := io.MultiWriter(config.Outputs...)

	// Create HandlerOptions with the desired log level.
	handlerOptions := &slog.HandlerOptions{
		Level: convertLogLevel(config.LogLevel),
	}

	// Define handler based on log format.
	var handler slog.Handler
	if config.OutputFormat == log.OutputFormatJSON {
		handler = slog.NewJSONHandler(multiWriter, handlerOptions)
	} else {
		handler = slog.NewTextHandler(multiWriter, handlerOptions)
	}

	return &SlogLogger{
		logger:   slog.New(handler),
		logLevel: config.LogLevel,
	}
}

func (l *SlogLogger) Debug(ctx context.Context, msg string, keysAndValues ...any) {
	if l.logLevel <= log.Debug {
		l.logger.DebugContext(ctx, msg, keysAndValues...)
	}
}

func (l *SlogLogger) Info(ctx context.Context, msg string, keysAndValues ...any) {
	if l.logLevel <= log.Info {
		l.logger.InfoContext(ctx, msg, keysAndValues...)
	}
}

func (l *SlogLogger) Warn(ctx context.Context, msg string, keysAndValues ...any) {
	if l.logLevel <= log.Warn {
		l.logger.WarnContext(ctx, msg, keysAndValues...)
	}
}

func (l *SlogLogger) Error(ctx context.Context, msg string, keysAndValues ...any) {
	if l.logLevel <= log.Error {
		l.logger.ErrorContext(ctx, msg, keysAndValues...)
	}
}

// Fatal always logs to ErrorContext irrespective of log level and calls os.Exit(1).
func (l *SlogLogger) Fatal(ctx context.Context, msg string, keysAndValues ...any) {
	l.logger.ErrorContext(ctx, msg, keysAndValues...)
	os.Exit(1)
}
