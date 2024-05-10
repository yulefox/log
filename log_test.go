package log

import (
	"sync"
	"testing"
)

func TestPrepare(t *testing.T) {
	logger := getDefaultLogger("default")
	if logger == nil {
		t.Errorf("Prepare failed, expected logger to be set, got nil")
	}
}

func TestLogger(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		Info("test", "This is a named logger ", "hello ", i, i+1, i == 2)
		go func() {
			logger, err := GetLogger("named", func(options *Options) error {
				return nil
			})
			if err != nil {
				t.Errorf("Failed to get named logger")
			}
			logger.Error("This is a named logger ", i, i+1)
			Error("test", "This is a named logger ", i)
			logger.Info("This is a named logger ", i)
			logger.Info("This is a named logger")
			logger.Warn("This is a logger: %s", "named")
			logger.Error("This is a named logger")
			Info("This is a named logger")
			Warn("This is a logger: %s", "named")
			Error("This is a named logger")
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestInfo(t *testing.T) {
	Info("%s %s", "param1", "param2")
	// Add assertions based on expected behavior
}

func TestWarn(t *testing.T) {
	Warn("param1", "param2")
	// Add assertions based on expected behavior
}

func TestError(t *testing.T) {
	Error("%s %s", "param1", "param2")
	// Add assertions based on expected behavior
}

func TestPanic(t *testing.T) {
	// Use defer and recover to catch the panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Panic did not cause a panic")
		}
	}()
	Panic("%s %s", "param1", "param2")
}

func TestFatal(t *testing.T) {
	// Fatal is expected to call os.Exit, testing it may require special handling
	Fatal("%s %s", "param1", "param2")
}

func BenchmarkDebug(b *testing.B) {
	// setup
	logger := getDefaultLogger("default")

	// run the Debug function b.N times
	for i := 0; i < b.N; i++ {
		logger.Debug("%s %s", "param1", "param2")
	}
}

func BenchmarkError(b *testing.B) {
	// setup
	logger := getDefaultLogger("default")

	// run the Debug function b.N times
	for i := 0; i < b.N; i++ {
		logger.Error("%s %s", "param1", "param2")
	}
}
