package caller

import (
	"fmt"
	"testing"
)

func TestGetCallerInfo(t *testing.T) {
	callerInfo, err := GetCaller(0)
	if err != nil {
		t.Fatalf("GetCallerInfo failed: %v", err)
	}

	if callerInfo.File == "" || callerInfo.Line == 0 || callerInfo.Function == "" {
		t.Errorf("Expected valid caller info, got: %+v", callerInfo)
	}

	fmt.Printf("CallerInfo: %+v\n", callerInfo)
}

func TestGetStackTrace(t *testing.T) {
	stackTrace, err := GetStackTrace(0)
	if err != nil {
		t.Fatalf("GetStackTrace failed: %v", err)
	}

	if len(stackTrace) == 0 {
		t.Error("Expected non-empty stack trace")
	}

	fmt.Printf("StackTrace:\n%s", stackTrace.String())
}

func TestGetCallerInfoWithSkip(t *testing.T) {
	callerInfo, err := GetCaller(1)
	if err != nil {
		t.Fatalf("GetCallerInfo failed: %v", err)
	}

	if callerInfo.File == "" || callerInfo.Line == 0 || callerInfo.Function == "" {
		t.Errorf("Expected valid caller info, got: %+v", callerInfo)
	}

	fmt.Printf("CallerInfo with skip: %+v\n", callerInfo)
}

func TestGetStackTraceWithSkip(t *testing.T) {
	stackTrace, err := GetStackTrace(1)
	if err != nil {
		t.Fatalf("GetStackTrace failed: %v", err)
	}

	if len(stackTrace) == 0 {
		t.Error("Expected non-empty stack trace")
	}

	fmt.Printf("StackTrace with skip:\n%s", stackTrace.String())
}
