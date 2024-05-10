package log

import (
	"fmt"
	"os"
	"runtime"
	"sync/atomic"
	"time"
)

const (
	Version = "0.1.0"
)

var (
	_logger atomic.Pointer[Logger]
)

type Logger struct {
	Options Options `json:"options"`
}

func getLogger(options ...Option) *Logger {
	name := os.Getenv("APP_NAME")
	if logger := _logger.Load(); logger != nil {
		return logger
	}

	opts := GetDefaultOptions(name)
	for _, opt := range options {
		if opt != nil {
			if err := opt(&opts); err != nil {
				return nil
			}
		}
	}

	logger := &Logger{
		Options: opts,
	}
	_logger.Store(logger)
	return logger
}

func (l *Logger) do(level Level, tag string, params ...any) {
	if level < l.Options.Level {
		return
	}

	ac := &Action{
		Level:   level,
		Options: &l.Options,
	}

	if l.Options.TimeFormat != "" {
		ac.Date = time.Now().Format(l.Options.TimeFormat)
	}

	if l.Options.AddCaller || ac.Level >= ERRO {
		var pc [10]uintptr
		n := runtime.Callers(l.Options.Skip, pc[:])
		if n > 0 {
			callers := pc[:n]
			frames := runtime.CallersFrames(callers)
			frame, more := frames.Next()
			caller := func() string {
				return fmt.Sprintf("%v:%v %v",
					TrimPath(frame.File),
					frame.Line,
					frame.Function,
				)
			}
			if l.Options.AddCaller {
				ac.Caller = fmt.Sprintf("%v:%v",
					TrimPath(frame.File),
					frame.Line)
			}

			switch ac.Level {
			case ERRO, PNIC, FATL:
				for more {
					ac.Stack = append(ac.Stack, caller())
					frame, more = frames.Next()
				}
			default:
			}
		}
	}
	if level == FATL {
		ac.AfterWrite = func() {
			os.Exit(1)
		}
	}

	ac.Name = tag
	ac.Write(params...)
}

func Debug(tag string, params ...any) {
	if l := getLogger(); l != nil {
		l.do(DEBU, tag, params...)
	}
}

func Info(tag string, params ...any) {
	if l := getLogger(); l != nil {
		l.do(INFO, tag, params...)
	}
}

func Warn(tag string, params ...any) {
	if l := getLogger(); l != nil {
		l.do(WARN, tag, params...)
	}
}

func Error(tag string, params ...any) {
	if l := getLogger(); l != nil {
		l.do(ERRO, tag, params...)
	}
}

func Panic(tag string, params ...any) {
	if l := getLogger(); l != nil {
		l.do(PNIC, tag, params...)
	}
}

func Fatal(tag string, params ...any) {
	if l := getLogger(); l != nil {
		l.do(FATL, tag, params...)
	}
}
