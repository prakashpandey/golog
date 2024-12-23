package zap

import (
	"context"

	"github.com/prakashpandey/golog/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogger is an implementation of Logger interface using Uber's Zap.
type ZapLogger struct {
	logger   *zap.Logger
	logLevel log.Level
}

// convertLogLevel converts a custom log.Level to zapcore.Level.
func convertLogLevel(level log.Level) zapcore.Level {
	var l zapcore.Level
	switch level {
	case log.Debug:
		l = zapcore.DebugLevel
	case log.Info:
		l = zapcore.InfoLevel
	case log.Warn:
		l = zapcore.WarnLevel
	case log.Error:
		l = zapcore.ErrorLevel
	default:
		l = zapcore.InfoLevel
	}
	return l
}

// NewZapLogger creates a new instance of ZapLogger with the given configuration.
func NewZapLogger(config log.Config) log.Logger {
	zapConfig := zap.NewProductionConfig()
	var cores []zapcore.Core
	for _, output := range config.Outputs {
		writer := zapcore.AddSync(output)
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(zapConfig.EncoderConfig),
			writer,
			convertLogLevel(config.LogLevel),
		)
		cores = append(cores, core)
	}
	logger := zap.New(zapcore.NewTee(cores...))

	return &ZapLogger{
		logger:   logger,
		logLevel: config.LogLevel,
	}

}

// convertToZapFields converts keysAndValues to zap.Fields.
func convertToZapFields(keysAndValues ...any) []zap.Field {
	fields := make([]zap.Field, 0, len(keysAndValues)/2)
	for i := 0; i < len(keysAndValues); i += 2 {
		if i+1 < len(keysAndValues) {
			key, ok := keysAndValues[i].(string)
			if !ok {
				continue
			}
			fields = append(fields, zap.Any(key, keysAndValues[i+1]))
		}
	}
	return fields
}

// Debug logs a message at DebugLevel.
func (l *ZapLogger) Debug(ctx context.Context, msg string, keysAndValues ...any) {
	if l.logLevel <= log.Debug {
		l.logger.Debug(msg, convertToZapFields(keysAndValues...)...)
	}
}

// Info logs a message at InfoLevel.
func (l *ZapLogger) Info(ctx context.Context, msg string, keysAndValues ...any) {
	if l.logLevel <= log.Info {
		l.logger.Info(msg, convertToZapFields(keysAndValues...)...)
	}
}

// Warn logs a message at WarnLevel.
func (l *ZapLogger) Warn(ctx context.Context, msg string, keysAndValues ...any) {
	if l.logLevel <= log.Warn {
		l.logger.Warn(msg, convertToZapFields(keysAndValues...)...)
	}
}

// Error logs a message at ErrorLevel.
func (l *ZapLogger) Error(ctx context.Context, msg string, keysAndValues ...any) {
	if l.logLevel <= log.Error {
		l.logger.Error(msg, convertToZapFields(keysAndValues...)...)
	}
}

// Fatal logs a message at ErrorLevel and then calls os.Exit(1).
func (l *ZapLogger) Fatal(ctx context.Context, msg string, keysAndValues ...any) {
	l.logger.Fatal(msg, convertToZapFields(keysAndValues...)...)
}
