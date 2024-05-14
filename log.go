package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync/atomic"
	"time"
)

const (
	Version = "0.1.0"
	AppName = "APP_NAME"
	AppMode = "APP_MODE"
)

var (
	_logger atomic.Pointer[Logger]
)

type Logger struct {
	Options Options `json:"options"`
}

// SetCaller Set whether to show caller
func SetCaller(show bool) Option {
	return func(o *Options) error {
		o.AddCaller = show
		return nil
	}
}

// SetLevel Set log level
func SetLevel(level Level) Option {
	return func(o *Options) error {
		o.Level = level
		return nil
	}
}

// SetTimeFormat Set time format
func SetTimeFormat(format string, nowFunc func() time.Time) Option {
	return func(o *Options) error {
		o.TimeFormat = format
		o.Now = nowFunc
		return nil
	}
}

// AddFileLogger Add file logger
func AddFileLogger(name string, writers ...io.Writer) Option {
	return func(o *Options) error {
		o.Cores = append(o.Cores, NewFileCore(name, writers...))
		return nil
	}
}

// AddJsonLogger Add json logger
func AddJsonLogger(writer io.Writer) Option {
	return func(o *Options) error {
		o.Cores = append(o.Cores, NewJsonCore(writer))
		return nil
	}
}

func Init(options ...Option) *Logger {
	opts := GetDefaultOptions()
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

func getLogger(options ...Option) *Logger {
	if logger := _logger.Load(); logger != nil {
		return logger
	}

	return Init(options...)
}

func (o Options) do(level Level, tag string, params ...any) {
	if level < o.Level {
		return
	}

	ac := &Action{
		Level:   level,
		Options: &o,
	}

	if o.TimeFormat != "" {
		ac.Date = o.Now().Format(o.TimeFormat)
	}

	if o.AddCaller || ac.Level >= ERRO {
		var pc [10]uintptr
		n := runtime.Callers(o.Skip, pc[:])
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
			if o.AddCaller {
				ac.Caller = fmt.Sprintf("%v:%v",
					TrimPath(frame.File),
					frame.Line)
			}

			addFrames := func() {
				for more {
					ac.Stack = append(ac.Stack, caller())
					frame, more = frames.Next()
				}
			}
			switch ac.Level {
			case ERRO, FATL, PNIC:
				addFrames()
			default:
			}
		}
	}

	ac.Tag = tag
	ac.Write(params...)
}

func Debug(tag string, params ...any) {
	if l := getLogger(); l != nil {
		l.Options.do(DEBU, tag, params...)
	}
}

func Info(tag string, params ...any) {
	if l := getLogger(); l != nil {
		l.Options.do(INFO, tag, params...)
	}
}

func Warn(tag string, params ...any) {
	if l := getLogger(); l != nil {
		l.Options.do(WARN, tag, params...)
	}
}

func Error(tag string, params ...any) {
	if l := getLogger(); l != nil {
		l.Options.do(ERRO, tag, params...)
	}
}

func Panic(tag string, params ...any) {
	if l := getLogger(); l != nil {
		l.Options.do(PNIC, tag, params...)
	}
	panic(params)
}

func Fatal(tag string, params ...any) {
	if l := getLogger(); l != nil {
		l.Options.do(FATL, tag, params...)
	}
	os.Exit(1)
}
