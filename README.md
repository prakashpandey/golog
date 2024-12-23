# GoLog Logger Module

GoLog is a Go logging module that provides an interface-based logging system with support for configurable log levels and multiple output targets. It can integrate seamlessly with any logging library, such as `log/slog` from the standard golang package and `Uber Zap`, both of which are already implemented as examples. Additional libraries can be integrated in a similar manner.


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
		Outputs:      []io.Writer{os.Stdout, os.Stderr},
		OutputFormat: log.OutputFormatTEXT,
		LogLevel:     log.Info,
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
		Outputs:      []io.Writer{os.Stdout, os.Stderr},
		OutputFormat: log.OutputFormatJSON,
		LogLevel:     log.Info,
	}

	logger := zap.NewZapLogger(config)
	ctx := context.Background()

	logger.Info(ctx, "Application started", "version", "1.0.0")
	logger.Warn(ctx, "This is a warning message", "component", "main")
	logger.Error(ctx, "An error occurred", "error", "nil pointer dereference")
}
```
