package log

import "time"

const (
	DefaultName       = ""
	DefaultLevel      = DEBU
	DefaultSkip       = 3
	DefaultAddCaller  = true
	DefaultTimeFormat = "2006-01-02 15:04:05.000"
)

// Option is a function on the options for a logger.
type Option func(*Options) error

// Options can be used to create a customized logger.
type Options struct {
	// TimeFormat is the time format for log entries.
	TimeFormat string

	// Level is the log level the logger should log at.
	Level Level

	// AddCaller is a flag to add the caller to the log entry.
	AddCaller bool

	// Skip is the number of frames to skip when computing the file name and line number.
	Skip int

	// Cores is a list of Cores the logger should write to.
	Cores []*Core

	Now func() time.Time
}

// GetDefaultOptions returns default configuration options for the client.
func GetDefaultOptions() Options {
	return Options{
		Level:      DefaultLevel,
		TimeFormat: DefaultTimeFormat,
		AddCaller:  DefaultAddCaller,
		Skip:       DefaultSkip,
		Now:        time.Now,
		Cores:      []*Core{NewTermCore()},
	}
}
