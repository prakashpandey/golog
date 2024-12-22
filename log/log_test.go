package log

import (
	"testing"
)

func TestParseLogLevel(t *testing.T) {
	tests := []struct {
		input    string
		expected Level
		err      bool
	}{
		{"DEBUG", Debug, false},
		{"INFO", Info, false},
		{"WARN", Warn, false},
		{"ERROR", Error, false},
		{"INVALID", "", true},
		{"", "", true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := ParseLogLevel(test.input)
			if (err != nil) != test.err {
				t.Errorf("ParseLogLevel(%q) error = %v, expected error = %v", test.input, err, test.err)
			}
			if result != test.expected {
				t.Errorf("ParseLogLevel(%q) = %v, expected %v", test.input, result, test.expected)
			}
		})
	}
}

func TestParseLogLevel_ValidLevels(t *testing.T) {
	validLevels := []struct {
		input    string
		expected Level
	}{
		{"DEBUG", Debug},
		{"INFO", Info},
		{"WARN", Warn},
		{"ERROR", Error},
	}

	for _, level := range validLevels {
		t.Run(level.input, func(t *testing.T) {
			result, err := ParseLogLevel(level.input)
			if err != nil {
				t.Errorf("ParseLogLevel(%q) returned an unexpected error: %v", level.input, err)
			}
			if result != level.expected {
				t.Errorf("ParseLogLevel(%q) = %v, expected %v", level.input, result, level.expected)
			}
		})
	}
}

func TestParseLogLevel_InvalidLevels(t *testing.T) {
	invalidLevels := []string{
		"INVALID",
		"debug",
		"info",
		"warn",
		"error",
		"",
		"123",
		"DEBUG ",
		" INFO",
	}

	for _, level := range invalidLevels {
		t.Run(level, func(t *testing.T) {
			_, err := ParseLogLevel(level)
			if err == nil {
				t.Errorf("ParseLogLevel(%q) expected an error but got none", level)
			}
		})
	}
}
