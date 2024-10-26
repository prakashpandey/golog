package log

import (
	"context"
	"io"
)

// Level defines different levels of logging.
type Level int

const (
	Debug Level = iota
	Info
	Warn
	Error
	Critical
	Fatal
)

type Logger interface {
	Debug(ctx context.Context, msg string, keysAndValues ...any)
	Info(ctx context.Context, msg string, keysAndValues ...any)
	Warn(ctx context.Context, msg string, keysAndValues ...any)
	Error(ctx context.Context, msg string, keysAndValues ...any)
	Critical(ctx context.Context, msg string, keysAndValues ...any)
	Fatal(ctx context.Context, msg string, keysAndValues ...any)
}

// Conf holds configuration for the logger, including log level and output format.
type Conf struct {
	Outputs  []io.Writer // Output targets, e.g., os.Stdout, os.Stderr
	UseJSON  bool        // Set true for JSON output
	LogLevel Level       // Minimum log level
}
