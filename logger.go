package log

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/yulefox/log/core"
)

const (
	Version = "0.1.0"
)

var (
	_loggers sync.Map
)

type Logger struct {
	Options core.Options `json:"options"`
}

func GetLogger(name string, options ...core.Option) (*Logger, error) {
	if logger, ok := _loggers.Load(name); ok {
		return logger.(*Logger), nil
	}

	opts := core.GetDefaultOptions(name)
	for _, opt := range options {
		if opt != nil {
			if err := opt(&opts); err != nil {
				return nil, err
			}
		}
	}

	logger := &Logger{
		Options: opts,
	}
	_loggers.Store(name, logger)
	return logger, nil
}

func getDefaultLogger(name string) *Logger {
	logger, _ := GetLogger("default", func(options *core.Options) error {
		options.Skip = core.DefaultSkip + 1
		options.Name = name
		return nil
	})
	return logger
}

func Debug(name string, params ...any) {
	if logger := getDefaultLogger(name); logger != nil {
		logger.Debug(params...)
	}
}

func Info(name string, params ...any) {
	if logger := getDefaultLogger(name); logger != nil {
		logger.Info(params...)
	}
}

func Warn(name string, params ...any) {
	if logger := getDefaultLogger(name); logger != nil {
		logger.Warn(params...)
	}
}

func Error(name string, params ...any) {
	if logger := getDefaultLogger(name); logger != nil {
		logger.Error(params...)
	}
}

func Panic(name string, params ...any) {
	if logger := getDefaultLogger(name); logger != nil {
		logger.Panic(params...)
	}
}

func Fatal(name string, params ...any) {
	if logger := getDefaultLogger(name); logger != nil {
		logger.Fatal(params...)
	}
}

func (l *Logger) Debug(params ...any) {
	if ac := l.check(core.DEBU); ac != nil {
		ac.Write(params...)
	}
}

func (l *Logger) Info(params ...any) {
	if ac := l.check(core.INFO); ac != nil {
		ac.Write(params...)
	}
}

func (l *Logger) Warn(params ...any) {
	if ac := l.check(core.WARN); ac != nil {
		ac.Write(params...)
	}
}

func (l *Logger) Error(params ...any) {
	if ac := l.check(core.ERRO); ac != nil {
		ac.Write(params...)
	}
}

func (l *Logger) Panic(params ...any) {
	if ac := l.check(core.PNIC); ac != nil {
		ac.Write(params...)
	}
}

func (l *Logger) Fatal(params ...any) {
	if ac := l.check(core.FATL); ac != nil {
		ac.Write(params...)
	}
}

func (l *Logger) check(level core.Level) *core.Action {
	if level < l.Options.Level {
		return nil
	}

	ac := &core.Action{
		Level:   level,
		Options: &l.Options,
	}

	if l.Options.TimeFormat != "" {
		ac.Date = time.Now().Format(l.Options.TimeFormat)
	}

	if l.Options.AddCaller || ac.Level >= core.ERRO {
		var pc [10]uintptr
		n := runtime.Callers(l.Options.Skip, pc[:])
		if n > 0 {
			callers := pc[:n]
			frames := runtime.CallersFrames(callers)
			frame, more := frames.Next()
			caller := func() string {
				return fmt.Sprintf("%v:%v %v",
					core.TrimPath(frame.File),
					frame.Line,
					frame.Function,
				)
			}
			if l.Options.AddCaller {
				ac.Caller = fmt.Sprintf("%v:%v",
					core.TrimPath(frame.File),
					frame.Line)
			}

			switch ac.Level {
			case core.ERRO, core.PNIC, core.FATL:
				for more {
					ac.Stack = append(ac.Stack, caller())
					frame, more = frames.Next()
				}
			default:
			}
		}
	}
	if level == core.FATL {
		ac.AfterWrite = func() {
			os.Exit(1)
		}
	}

	return ac
}
