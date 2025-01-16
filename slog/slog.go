package slog

import (
	"context"
	"io"
	"log/slog"
	"os"

	"github.com/prakashpandey/golog/caller"
	"github.com/prakashpandey/golog/log"
)

// SlogLogger is a concrete implementation of the Logger interface using slog.
type SlogLogger struct {
	logger *slog.Logger
	log.Config
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
	config.Sanitize()
	config.Default()
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

	var attrs []slog.Attr
	for k, v := range config.Attrs {
		attrs = append(attrs, slog.Attr{
			Key:   k,
			Value: slog.AnyValue(v),
		})
	}

	handler = handler.WithAttrs(attrs)

	return &SlogLogger{
		logger: slog.New(handler),
		Config: config,
	}
}

// stacktrace returns the keys and values with caller and stack trace information.
// It appends caller and stack trace information to the keys and values if enabled.
// Caller information is appended only if enabled.
// Stack trace information is appended only if enabled and the log level is greater than or equal to the stack trace level.
func (l *SlogLogger) stacktrace(level log.Level, keysAndValues []any) []any {
	return caller.AddStacktrace(level, l.Config, keysAndValues)
}

func (l *SlogLogger) Debug(ctx context.Context, msg string, keysAndValues ...any) {
	if l.LogLevel <= log.Debug {
		l.logger.DebugContext(ctx, msg, l.stacktrace(log.Debug, keysAndValues)...)
	}
}

func (l *SlogLogger) Info(ctx context.Context, msg string, keysAndValues ...any) {
	if l.LogLevel <= log.Info {
		l.logger.InfoContext(ctx, msg, l.stacktrace(log.Info, keysAndValues)...)
	}
}

func (l *SlogLogger) Warn(ctx context.Context, msg string, keysAndValues ...any) {
	if l.LogLevel <= log.Warn {
		l.logger.WarnContext(ctx, msg, l.stacktrace(log.Warn, keysAndValues)...)
	}
}

func (l *SlogLogger) Error(ctx context.Context, msg string, keysAndValues ...any) {
	if l.LogLevel <= log.Error {
		l.logger.ErrorContext(ctx, msg, l.stacktrace(log.Error, keysAndValues)...)
	}
}

// Fatal always logs to ErrorContext irrespective of log level and calls os.Exit(1).
func (l *SlogLogger) Fatal(ctx context.Context, msg string, keysAndValues ...any) {
	l.logger.ErrorContext(ctx, msg, l.stacktrace(log.Error, keysAndValues)...)
	os.Exit(1)
}
