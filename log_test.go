package log

import (
	"bytes"
	"strings"
	"testing"
)

func TestInit(t *testing.T) {
	logger := Init(SetLevel(INFO))
	if logger.Options.Level != INFO {
		t.Errorf("Expected level to be INFO, got %v", logger.Options.Level)
	}

	logger = Init(SetCaller(true))
	if !logger.Options.AddCaller {
		t.Errorf("Expected AddCaller to be true, got %v", logger.Options.AddCaller)
	}
}

func TestGetLogger(t *testing.T) {
	logger := getLogger()
	if logger == nil {
		t.Error("Expected logger to be not nil")
	}
}

func TestDo(t *testing.T) {
	var buf bytes.Buffer
	logger := Init(SetLevel(INFO), AddJsonLogger(&buf))

	logger.Options.do(INFO, "test", "message")
	if !strings.Contains(buf.String(), "message") {
		t.Errorf("Expected 'message' to be in log output")
	}

	buf.Reset()
	logger.Options.do(DEBU, "test", "message")
	if buf.Len() != 0 {
		t.Errorf("Expected no log output for DEBUG level")
	}
}

func TestLogFunctions(t *testing.T) {
	var buf bytes.Buffer
	Init(SetLevel(INFO), AddJsonLogger(&buf))

	Debug("test", "debug message")
	if buf.Len() != 0 {
		t.Errorf("Expected no log output for DEBUG level")
	}

	Info("test", "info message")
	if !strings.Contains(buf.String(), "info message") {
		t.Errorf("Expected 'info message' to be in log output")
	}

	buf.Reset()
	Warn("test", "warn message")
	if !strings.Contains(buf.String(), "warn message") {
		t.Errorf("Expected 'warn message' to be in log output")
	}

	buf.Reset()
	Error("test", "error message")
	if !strings.Contains(buf.String(), "error message") {
		t.Errorf("Expected 'error message' to be in log output")
	}

	defer func() {
		if !strings.Contains(buf.String(), "panic message") {
			t.Errorf("Expected 'panic message' to be in log output")
		}
		if err := recover(); err != nil {
			//Fatal("test", "fatal message")
		}
	}()
	buf.Reset()
	Panic("panic", "panic message")
}

func BenchmarkError(b *testing.B) {
	// run the Debug function b.N times
	for i := 0; i < b.N; i++ {
		Info("", "%s %s", "param1", "param2")
		Error("", "%s %s", "param1", "param2")
	}
}
