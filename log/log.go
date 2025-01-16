package log

import (
	"context"
	"errors"
	"io"
	"os"
	"strings"
	"time"
)

// Level defines different levels of logging.
type Level int

const (
	Debug Level = iota
	Info
	Warn
	Error
)

func ParseLevel(levelStr string) (Level, error) {
	switch levelStr {
	case "DEBUG", "debug":
		return Debug, nil
	case "INFO", "info":
		return Info, nil
	case "WARN", "warn":
		return Warn, nil
	case "ERROR", "error":
		return Error, nil
	default:
		return 0, errors.New("invalid log level: " + levelStr)
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

type Caller struct {
	FieldName string // Key name for caller in the log. Default is "caller
	Enabled   bool
	Skip      int
}

type Stacktrace struct {
	FieldName string // Key name for stack trace in the log. Default is "stacktrace"
	Enabled   bool
	Level     Level // Minimum log level to record stack trace
}

// Config holds configuration for the logger, including log level and output format.
type Config struct {
	TmFn         func() time.Time  // Time function
	Caller       Caller            // Caller configuration
	Stacktrace   Stacktrace        // Stacktrace configuration
	Outputs      []io.Writer       // Output targets, e.g., os.Stdout, os.Stderr
	OutputFormat OutputFormat      // Output format
	LogLevel     Level             // Minimum log level
	Attrs        map[string]string // Additional attributes to be logged for each log entry.
}

func DefaultConfig() Config {
	return Config{
		TmFn:         time.Now,
		Outputs:      []io.Writer{os.Stdout},
		LogLevel:     Info,
		OutputFormat: OutputFormatTEXT,
		Caller:       Caller{Enabled: true, Skip: 1},
		Stacktrace:   Stacktrace{Enabled: true, Level: Error},
	}
}

// Sanitize trims leading and trailing white spaces from the field of type string.
func (c *Config) Sanitize() {
	c.Caller.FieldName = strings.TrimSpace(c.Caller.FieldName)
	c.Stacktrace.FieldName = strings.TrimSpace(c.Stacktrace.FieldName)
}

// Default sets default values for the logger configuration.
// If a field is already set, it will not be overridden.
// If a field is not set, it will be set to a default value.
// The default values are:
// - TmFn: time.Now
// - Caller.FieldName: "caller"
// - Stacktrace.FieldName: "stacktrace"
// - OutputFormat: OutputFormatTEXT
// - LogLevel: Info
func (c *Config) Default() {
	if c.TmFn == nil {
		c.TmFn = time.Now
	}
	if c.Caller.FieldName == "" {
		c.Caller.FieldName = "caller"
	}
	if c.Stacktrace.FieldName == "" {
		c.Stacktrace.FieldName = "stacktrace"
	}
	if c.OutputFormat == "" {
		c.OutputFormat = OutputFormatTEXT
	}
	if c.LogLevel == 0 {
		c.LogLevel = Info
	}
}
