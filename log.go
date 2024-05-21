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
	_logger       atomic.Pointer[Logger]
	_namedLoggers sync.Map
	_entryPool    sync.Pool
)

type Logger struct {
	options Options
}

// SetName Set the name of logger
func SetName(name string) Option {
	return func(o *Options) error {
		o.Name = name
		return nil
	}
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
		if name == "" && o.Name != "" {
			name = o.Name
		}
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
		options: opts,
	}
	if opts.Name != "" {
		_namedLoggers.Store(opts.Name, logger)
	} else {
		_logger.Store(logger)
	}
	return logger
}

func GetLogger(name string) *Logger {
	if logger, ok := _namedLoggers.Load(name); ok {
		return logger.(*Logger)
	}
	if logger := _logger.Load(); logger != nil {
		return logger
	}
	return Init(SetName(name))
}

func getDefaultLogger() *Logger {
	if logger := _logger.Load(); logger != nil {
		return logger
	}
	return Init()
}

func getEntry(options *Options) *Entry {
	entry, ok := _entryPool.Get().(*Entry)
	if ok {
		entry.Options = options
		return entry
	}
	return &Entry{
		Options: options,
	}
}

func putEntry(e *Entry) {
	e.Fields = []string{}
	e.Stack = []string{}
	e.AfterWrite = nil
	_entryPool.Put(e)
}

func (l *Logger) Debug(params ...any) {
	if e := getEntry(&l.options); e != nil {
		defer putEntry(e)
		e.log(DEBU, params...)
	}
}

func (l *Logger) Info(params ...any) {
	if e := getEntry(&l.options); e != nil {
		defer putEntry(e)
		e.log(INFO, params...)
	}
}

func (l *Logger) Warn(params ...any) {
	if e := getEntry(&l.options); e != nil {
		defer putEntry(e)
		e.log(WARN, params...)
	}
}

func (l *Logger) Error(params ...any) {
	if e := getEntry(&l.options); e != nil {
		defer putEntry(e)
		e.log(ERRO, params...)
	}
}

func (l *Logger) Panic(params ...any) {
	if e := getEntry(&l.options); e != nil {
		defer putEntry(e)
		e.log(PNIC, params...)
	}
	panic(params)
}

func (l *Logger) Fatal(params ...any) {
	if e := getEntry(&l.options); e != nil {
		defer putEntry(e)
		e.log(FATL, params...)
	}
	os.Exit(1)
}

func Debug(params ...any) {
	l := getDefaultLogger()
	if l != nil {
		if e := getEntry(&l.options); e != nil {
			defer putEntry(e)
			e.log(DEBU, params...)
		}
	}
}

func Info(params ...any) {
	l := getDefaultLogger()
	if l != nil {
		if e := getEntry(&l.options); e != nil {
			defer putEntry(e)
			e.log(INFO, params...)
		}
	}
}

func Warn(params ...any) {
	l := getDefaultLogger()
	if l != nil {
		if e := getEntry(&l.options); e != nil {
			defer putEntry(e)
			e.log(WARN, params...)
		}
	}
}

func Error(params ...any) {
	l := getDefaultLogger()
	if l != nil {
		if e := getEntry(&l.options); e != nil {
			defer putEntry(e)
			e.log(ERRO, params...)
		}
	}
}

func Panic(params ...any) {
	l := getDefaultLogger()
	if l != nil {
		if e := getEntry(&l.options); e != nil {
			defer putEntry(e)
			e.log(PNIC, params...)
		}
	}
	panic(params)
}

func Fatal(params ...any) {
	l := getDefaultLogger()
	if l != nil {
		if e := getEntry(&l.options); e != nil {
			defer putEntry(e)
			e.log(FATL, params...)
		}
	}
	os.Exit(1)
}
