package main

import (
	"context"
	"io"
	"os"

	"github.com/prakashpandey/golog/log"
	"github.com/prakashpandey/golog/slog"
)

func main() {
	config := log.Config{
		Outputs:      []io.Writer{os.Stdout, os.Stderr},
		OutputFormat: log.OutputFormatJSON,
		LogLevel:     log.Info,
	}

	logger := slog.NewSlogLogger(config)
	ctx := context.Background()

	logger.Info(ctx, "Application started", "version", "1.0.0")
	logger.Warn(ctx, "This is a warning message", "component", "main")
	logger.Error(ctx, "An error occurred", "error", "nil pointer dereference")
}
