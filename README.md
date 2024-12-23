# GoLog Logger Module

`GoLog` is a Go logging module that implements an interface-based logging system with support for multiple output targets and configurable log levels. Internally, it uses the `slog` package for efficient logging.

## Features
- Supports multiple output targets (e.g., `stdout`, `stderr`).
- Supports both JSON and text log formats.
- Configurable log levels (Debug, Info, Warn, Error).
- Easily extendable for future logging backends.

## Installation

```sh
go get github.com/prakashpandey/golog
```

## Usage

Example 1: Using Slog

```golang
import(
	"github.com/prakashpandey/golog/log"
	"github.com/prakashpandey/golog/slog"
)
func main() {
	config := log.Conf{
		Outputs:  []io.Writer{os.Stdout, os.Stderr},
		UseJSON:  true,
		LogLevel: log.Info,
	}

	logger := slog.NewSlogLogger(config)
	ctx := context.Background()

	logger.Info(ctx, "Application started", "version", "1.0.0")
	logger.Warn(ctx, "This is a warning message", "component", "main")
	logger.Error(ctx, "An error occurred", "error", "nil pointer dereference")
}
```

Example 2: Using Uber Zap

```golang
import(
	"github.com/prakashpandey/golog/log"
	"github.com/prakashpandey/golog/zap"
)
func main() {
	config := log.Conf{
		Outputs:  []io.Writer{os.Stdout, os.Stderr},
		UseJSON:  true,
		LogLevel: log.Info,
	}

	logger := zap.NewZapLogger(config)
	ctx := context.Background()

	logger.Info(ctx, "Application started", "version", "1.0.0")
	logger.Warn(ctx, "This is a warning message", "component", "main")
	logger.Error(ctx, "An error occurred", "error", "nil pointer dereference")
}
```
