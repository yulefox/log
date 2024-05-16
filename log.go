package log

import (
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

const (
	AppName = "APP_NAME"
	AppMode = "APP_MODE"
)

var (
	_logger    atomic.Pointer[Logger]
	_entryPool sync.Pool
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

func getEntry(options ...Option) *Entry {
	logger := getLogger(options...)
	if logger == nil {
		return nil
	}
	entry, _ := _entryPool.Get().(*Entry)
	if entry == nil {
		entry = &Entry{
			Options: &logger.Options,
			channel: make(chan *logEntry, 1024),
		}
	}
	go func() {
		select {
		case data, closed := <-entry.channel:
			if !closed {
				putEntry(entry)
				return
			}
			entry._log(data)
			return
		}
	}()
	return entry
}

func putEntry(e *Entry) {
	e.Fields = []string{}
	e.Stack = []string{}
	e.channel = make(chan *logEntry, 1024)
	e.AfterWrite = nil
	_entryPool.Put(e)
}

func Debug(params ...any) {
	if e := getEntry(); e != nil {
		defer putEntry(e)
		e.log(DEBU, params...)
	}
}

func Info(params ...any) {
	if e := getEntry(); e != nil {
		defer putEntry(e)
		e.log(INFO, params...)
	}
}

func Warn(params ...any) {
	if e := getEntry(); e != nil {
		defer putEntry(e)
		e.log(WARN, params...)
	}
}

func Error(params ...any) {
	if e := getEntry(); e != nil {
		defer putEntry(e)
		e.log(ERRO, params...)
	}
}

func Panic(params ...any) {
	if e := getEntry(); e != nil {
		defer putEntry(e)
		e.log(PNIC, params...)
	}
	panic(params)
}

func Fatal(params ...any) {
	if e := getEntry(); e != nil {
		defer putEntry(e)
		e.log(FATL, params...)
	}
	os.Exit(1)
}
