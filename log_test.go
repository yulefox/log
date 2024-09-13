package log

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	var buf bytes.Buffer
	logger := Init(SetLevel(DEBU), AddFileLogger("", &buf))
	if logger.options.Level != DEBU {
		t.Errorf("Expected level to be INFO, got %v", logger.options.Level)
	}

	Debug("debug message")
	if !strings.Contains(buf.String(), "debug message") {
		t.Errorf("Expected 'debug message' to be in log output")
	}

	logger = Init(SetCaller(true), SetTimeFormat("2006-01-02 15:04:05.678", time.Now().UTC))
	if !logger.options.AddCaller {
		t.Errorf("Expected AddCaller to be true, got %v", logger.options.AddCaller)
	}

	logger = Init(func(options *Options) error {
		options.Cores = append(options.Cores, nil)
		return nil
	})
	if logger.options.Cores == nil || len(logger.options.Cores) == 0 || logger.options.Cores[len(logger.options.Cores)-1] != nil {
		t.Errorf("Expected Cores to contain a nil core, got %v", logger.options.Cores)
	}

	buf.Reset()
	Init(SetLevel(DEBU), AddJsonLogger(&buf), SetCaller(true))
	Info("info message")
	if !strings.Contains(buf.String(), "info message") {
		t.Errorf("Expected 'info message' to be in log output")
	}
}

func TestGetLogger(t *testing.T) {
	logger := getDefaultLogger()
	if logger == nil {
		t.Error("Expected logger to be not nil")
	}

	var buf bytes.Buffer
	Init(SetName("test"), SetLevel(INFO), AddFileLogger("", &buf))
	logger = GetLogger("test")
	if logger == nil {
		t.Error("Expected logger to be not nil")
	}
}

func TestNamedLoggerFunctions(t *testing.T) {
	var buf bytes.Buffer
	var bufError bytes.Buffer
	logger := Init(SetName("hello"), SetTimeFormat("2006-01-02 15:04:05.000000000", time.Now().UTC), SetLevel(INFO), AddFileLogger("test", &buf), AddFileLogger("test", &bufError))

	logger.Debug("debug message")
	if buf.Len() != 0 {
		t.Errorf("Expected no log output for DEBUG level")
	}

	logger.Info("info message: %s", "hello, world")
	if !strings.Contains(buf.String(), "info message: hello, world") {
		t.Errorf("Expected 'info message: hello, world' to be in log output, get %s", buf.String())
	}

	buf.Reset()
	logger.Warn("warn message")
	if !strings.Contains(buf.String(), "warn message") {
		t.Errorf("Expected 'warn message' to be in log output")
	}

	bufError.Reset()
	logger.Error("error message")
	if !strings.Contains(bufError.String(), "error message") {
		t.Errorf("Expected 'error message' to be in log output")
	}

	defer func() {
		if !strings.Contains(bufError.String(), "panic message") {
			t.Errorf("Expected 'panic message' to be in log output")
		}
		if err := recover(); err != nil {
			buf.Reset()
			logger.Warn("warn message")
			if !strings.Contains(buf.String(), "warn message") {
				t.Errorf("Expected 'warn message' to be in log output")
			}
		}
	}()
	bufError.Reset()
	logger.Panic("panic", "panic message")
}

func TestNamedFileLoggerFunctions(t *testing.T) {
	logger := Init(SetName("hello"), SetTimeFormat("2006-01-02 15:04:05.000000000", time.Now().UTC),
		SetCaller(true),
		SetLevel(INFO), AddFileLogger("test"),
		AddFileLogger("rpc"))

	logger.Debug("debug message")
	logger.Info("info message: %s", "hello, world")
	logger.Warn("warn message")
	logger.Error("error message")

	time.Sleep(5 * time.Second)
	defer func() {
		if err := recover(); err != nil {
			logger.Warn("warn message")
		}
		Fini()
	}()
	logger.Panic("panic message")
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

func TestFatal(t *testing.T) {
	Fatal("fatal message")
}

func TestAddFileLogger(t *testing.T) {
	Init(SetLevel(DEBU), AddFileLogger("test"))
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
