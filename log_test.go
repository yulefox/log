package log

import (
	"bytes"
	"os"
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

	logger = Init(func(options *Options) error {
		options.Cores = append(options.Cores, nil)
		return nil
	})
	if logger.Options.Cores == nil || len(logger.Options.Cores) == 0 || logger.Options.Cores[len(logger.Options.Cores)-1] != nil {
		t.Errorf("Expected Cores to contain a nil core, got %v", logger.Options.Cores)
	}
	var buf bytes.Buffer
	Init(SetLevel(INFO), AddJsonLogger(&buf), SetCaller(true))
	Info("info message")
	if !strings.Contains(buf.String(), "info message") {
		t.Errorf("Expected 'info message' to be in log output")
	}
}

func TestGetLogger(t *testing.T) {
	logger := getEntry()
	if logger == nil {
		t.Error("Expected logger to be not nil")
	}
}

func TestLogFunctions(t *testing.T) {
	var buf bytes.Buffer
	Init(SetLevel(INFO), AddFileLogger("test", &buf))

	Debug("debug message")
	if buf.Len() != 0 {
		t.Errorf("Expected no log output for DEBUG level")
	}

	Info("info message: %s", "hello, world")
	if !strings.Contains(buf.String(), "info message: hello, world") {
		t.Errorf("Expected 'info message: hello, world' to be in log output, get %s", buf.String())
	}

	buf.Reset()
	Warn("warn message")
	if !strings.Contains(buf.String(), "warn message") {
		t.Errorf("Expected 'warn message' to be in log output")
	}

	buf.Reset()
	Error("error message")
	if !strings.Contains(buf.String(), "error message") {
		t.Errorf("Expected 'error message' to be in log output")
	}

	defer func() {
		if !strings.Contains(buf.String(), "panic message") {
			t.Errorf("Expected 'panic message' to be in log output")
		}
		if err := recover(); err != nil {
			buf.Reset()
			Warn("warn message")
			if !strings.Contains(buf.String(), "warn message") {
				t.Errorf("Expected 'warn message' to be in log output")
			}
		}
	}()
	buf.Reset()
	Panic("panic", "panic message")
}

func TestAddFileLogger(t *testing.T) {
	Init(SetLevel(INFO), AddFileLogger("test"))
	Info("info message")

	filename := "logs/test.log"
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		t.Errorf("Expected '%s' to exist", filename)
	}
}

func BenchmarkError(b *testing.B) {
	// run the Debug function b.N times
	for i := 0; i < b.N; i++ {
		Info("", "%s %s", "param1", "param2")
		Error("", "%s %s", "param1", "param2")
	}
}
