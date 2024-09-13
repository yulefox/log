package log

import (
	"context"
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
	context.Context
	cancel  context.CancelFunc
	logs    chan *Entry
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

	ctx, cancel := context.WithCancel(context.Background())
	logger := &Logger{
		Context: ctx,
		cancel:  cancel,
		options: opts,
		logs:    make(chan *Entry, 1024),
	}
	if opts.Name != "" {
		_namedLoggers.Store(opts.Name, logger)
	} else {
		_logger.Store(logger)
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case e, ok := <-logger.logs:
				e.log()
				putEntry(e)
				if !ok {
					return
				}
			}
		}
	}()
	return logger
}

func Fini() {
	if logger := _logger.Load(); logger != nil {
		logger.cancel()
	}
	_namedLoggers.Range(func(key, value any) bool {
		value.(*Logger).cancel()
		return true
	})
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
	if l.options.Level > DEBU {
		return
	}
	if e := getEntry(&l.options); e != nil {
		l.log(e, DEBU, params...)
	}
}

func (l *Logger) Info(params ...any) {
	if l.options.Level > INFO {
		return
	}
	if e := getEntry(&l.options); e != nil {
		l.log(e, INFO, params...)
	}
}

func (l *Logger) Warn(params ...any) {
	if l.options.Level > WARN {
		return
	}
	if e := getEntry(&l.options); e != nil {
		l.log(e, WARN, params...)
	}
}

func (l *Logger) Error(params ...any) {
	if e := getEntry(&l.options); e != nil {
		l.log(e, ERRO, params...)
	}
}

func (l *Logger) Panic(params ...any) {
	if e := getEntry(&l.options); e != nil {
		l.log(e, PNIC, params...)
	}
	panic(params)
}

func (l *Logger) Fatal(params ...any) {
	if e := getEntry(&l.options); e != nil {
		l.log(e, FATL, params...)
	}
	os.Exit(1)
}

func (l *Logger) log(e *Entry, level Level, params ...any) {
	e.Level = level
	e.Params = params

	if e.TimeFormat != "" {
		e.Date = e.Now().Format(e.TimeFormat)
	}

	if e.AddCaller || e.Level >= ERRO {
		stack := GetStack(e.Skip, 10)
		if len(stack) > 0 {
			if e.AddCaller {
				e.Caller = e.FormatFrame(stack[0])
			}
			switch e.Level {
			case ERRO, FATL, PNIC:
				e.Stack = toStrings(stack)
			default:
			}
		}
	}
	select {
	case l.logs <- e:
	case <-time.After(1 * time.Second):
		return
	}
}

func Debug(params ...any) {
	l := getDefaultLogger()
	if l != nil {
		if l.options.Level > DEBU {
			return
		}
		if e := getEntry(&l.options); e != nil {
			l.log(e, DEBU, params...)
		}
	}
}

func Info(params ...any) {
	l := getDefaultLogger()
	if l != nil {
		if l.options.Level > INFO {
			return
		}
		if e := getEntry(&l.options); e != nil {
			l.log(e, INFO, params...)
		}
	}
}

func Warn(params ...any) {
	l := getDefaultLogger()
	if l != nil {
		if l.options.Level > WARN {
			return
		}
		if e := getEntry(&l.options); e != nil {
			l.log(e, WARN, params...)
		}
	}
}

func Error(params ...any) {
	l := getDefaultLogger()
	if l != nil {
		if e := getEntry(&l.options); e != nil {
			l.log(e, ERRO, params...)
		}
	}
}

func Panic(params ...any) {
	l := getDefaultLogger()
	if l != nil {
		if e := getEntry(&l.options); e != nil {
			l.log(e, PNIC, params...)
		}
	}
	panic(params)
}

func Fatal(params ...any) {
	l := getDefaultLogger()
	if l != nil {
		if e := getEntry(&l.options); e != nil {
			l.log(e, FATL, params...)
		}
	}
	os.Exit(1)
}
