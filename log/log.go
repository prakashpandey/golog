package log

import (
	"context"
	"errors"
	"io"
)

// Level defines different levels of logging.
type Level string

const (
	Debug Level = "DEBUG"
	Info  Level = "INFO"
	Warn  Level = "WARN"
	Error Level = "ERROR"
)

func ParseLogLevel(levelStr string) (Level, error) {
	switch levelStr {
	case string(Debug):
		return Debug, nil
	case string(Info):
		return Info, nil
	case string(Warn):
		return Warn, nil
	case string(Error):
		return Error, nil
	default:
		return "", errors.New("invalid log level: " + levelStr)
	}
}

type Logger interface {
	Debug(ctx context.Context, msg string, keysAndValues ...any)
	Info(ctx context.Context, msg string, keysAndValues ...any)
	Warn(ctx context.Context, msg string, keysAndValues ...any)
	Error(ctx context.Context, msg string, keysAndValues ...any)
	Fatal(ctx context.Context, msg string, keysAndValues ...any)
}

type OutputFormat string

const (
	OutputFormatTEXT OutputFormat = "TEXT"
	OutputFormatJSON OutputFormat = "JSON"
)

// Config holds configuration for the logger, including log level and output format.
type Config struct {
	Outputs      []io.Writer  // Output targets, e.g., os.Stdout, os.Stderr
	OutputFormat OutputFormat // Set true for JSON output
	LogLevel     Level        // Minimum log level
}
