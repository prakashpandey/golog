package zap

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/prakashpandey/golog/log"
	"go.uber.org/zap"
)

func TestNewZapLogger(t *testing.T) {
	config := log.Config{
		Outputs:      []io.Writer{&bytes.Buffer{}},
		OutputFormat: log.OutputFormatTEXT,
		LogLevel:     log.Debug,
	}

	logger := NewZapLogger(config)
	if logger == nil {
		t.Fatal("Expected logger to be non-nil")
	}
}

func TestZapLogger_Debug(t *testing.T) {
	var buf bytes.Buffer
	config := log.Config{
		Outputs:      []io.Writer{&buf},
		OutputFormat: log.OutputFormatTEXT,
		LogLevel:     log.Debug,
	}

	logger := NewZapLogger(config)
	ctx := context.Background()
	logger.Debug(ctx, "Debug message", "key", "value")

	if !bytes.Contains(buf.Bytes(), []byte("Debug message")) {
		t.Errorf("Expected 'Debug message' in log output, got: %s", buf.String())
	}
}

func TestZapLogger_Info(t *testing.T) {
	var buf bytes.Buffer
	config := log.Config{
		Outputs:      []io.Writer{&buf},
		OutputFormat: log.OutputFormatTEXT,
		LogLevel:     log.Info,
	}

	logger := NewZapLogger(config)
	ctx := context.Background()
	logger.Info(ctx, "Info message", "key", "value")

	if !bytes.Contains(buf.Bytes(), []byte("Info message")) {
		t.Errorf("Expected 'Info message' in log output, but it was not found")
	}
}

func TestZapLogger_Warn(t *testing.T) {
	var buf bytes.Buffer
	config := log.Config{
		Outputs:      []io.Writer{&buf},
		OutputFormat: log.OutputFormatTEXT,
		LogLevel:     log.Warn,
	}

	logger := NewZapLogger(config)
	ctx := context.Background()
	logger.Warn(ctx, "Warn message", "key", "value")

	if !bytes.Contains(buf.Bytes(), []byte("Warn message")) {
		t.Errorf("Expected 'Warn message' in log output, but it was not found")
	}
}

func TestZapLogger_Error(t *testing.T) {
	var buf bytes.Buffer
	config := log.Config{
		Outputs:      []io.Writer{&buf},
		OutputFormat: log.OutputFormatTEXT,
		LogLevel:     log.Error,
	}

	logger := NewZapLogger(config)
	ctx := context.Background()
	logger.Error(ctx, "Error message", "key", "value")

	if !bytes.Contains(buf.Bytes(), []byte("Error message")) {
		t.Errorf("Expected 'Error message' in log output, but it was not found")
	}
}

func TestConvertToZapFields(t *testing.T) {
	tests := []struct {
		name           string
		keysAndValues  []any
		expectedFields []zap.Field
	}{
		{
			name:           "Empty input",
			keysAndValues:  []any{},
			expectedFields: []zap.Field{},
		},
		{
			name: "Valid key-value pairs",
			keysAndValues: []any{
				"key1", "value1",
				"key2", 2,
				"key3", true,
			},
			expectedFields: []zap.Field{
				zap.Any("key1", "value1"),
				zap.Any("key2", 2),
				zap.Any("key3", true),
			},
		},
		{
			name: "Invalid key type",
			keysAndValues: []any{
				123, "value1",
				"key2", 2,
			},
			expectedFields: []zap.Field{
				zap.Any("key2", 2),
			},
		},
		{
			name: "Odd number of elements",
			keysAndValues: []any{
				"key1", "value1",
				"key2",
			},
			expectedFields: []zap.Field{
				zap.Any("key1", "value1"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := convertToZapFields(tt.keysAndValues...)
			if len(fields) != len(tt.expectedFields) {
				t.Errorf("Expected %d fields, got %d", len(tt.expectedFields), len(fields))
			}
			for i, field := range fields {
				if field.Key != tt.expectedFields[i].Key || field.Interface != tt.expectedFields[i].Interface {
					t.Errorf("Expected field %v, got %v", tt.expectedFields[i], field)
				}
			}
		})
	}
}
