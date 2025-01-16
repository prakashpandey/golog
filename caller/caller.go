package caller

import (
	"fmt"
	syslog "log"
	"runtime"
	"strings"

	"github.com/prakashpandey/golog/log"
)

// Caller represents the information about the caller.
type Caller struct {
	File     string
	Line     int
	Function string
}

// StackTrace represents a stack trace.
type StackTrace []Caller

// GetCaller returns the caller information for the given skip level.
func GetCaller(skip int) (Caller, error) {
	pc, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return Caller{}, fmt.Errorf("failed to get caller info")
	}

	funcName := runtime.FuncForPC(pc).Name()
	return Caller{
		File:     file,
		Line:     line,
		Function: funcName,
	}, nil
}

func (c Caller) String() string {
	return fmt.Sprintf("%s:%d %s", c.File, c.Line, c.Function)
}

// GetStackTrace returns the stack trace starting from the given skip level.
func GetStackTrace(skip int) (StackTrace, error) {
	var stackTrace StackTrace

	for i := skip + 1; ; i++ {
		callerInfo, err := GetCaller(i)
		if err != nil {
			break
		}
		stackTrace = append(stackTrace, callerInfo)
	}

	if len(stackTrace) == 0 {
		return nil, fmt.Errorf("failed to get stack trace")
	}

	return stackTrace, nil
}

// String returns a formatted string representation of the stack trace.
func (st StackTrace) String() string {
	var builder strings.Builder
	for _, callerInfo := range st {
		builder.WriteString(fmt.Sprintf("%s:%d %s\n", callerInfo.File, callerInfo.Line, callerInfo.Function))
	}
	return builder.String()
}

// stacktrace returns the keys and values with caller and stack trace information.
// It appends caller and stack trace information to the keys and values if enabled.
// Caller information is appended only if enabled.
// Stack trace information is appended only if enabled and the log level is greater than or equal to the stack trace level.
func AddStacktrace(level log.Level, config log.Config, keysAndValues []any) []any {
	if config.Caller.Enabled {
		c, err := GetCaller(config.Caller.Skip + 1)
		if err != nil {
			syslog.Default().Printf("failed to get caller info: %v", err)
		}
		// Append caller information to the keys and values at the start of the slice.
		keysAndValues = append([]any{config.Caller.FieldName, c.String()}, keysAndValues...)
	}

	if config.Stacktrace.Enabled && level >= config.Stacktrace.Level {
		st, err := GetStackTrace(config.Caller.Skip + 1)
		if err != nil {
			syslog.Default().Printf("failed to get stack trace: %v", err)
		}
		keysAndValues = append([]any{config.Stacktrace.FieldName, st.String()}, keysAndValues...)
	}

	return keysAndValues
}
