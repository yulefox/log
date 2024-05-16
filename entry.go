package log

import (
	"os"
	"runtime"
)

type Entry struct {
	*Options
	Level      Level
	Params     []any
	Date       string
	Caller     string
	Fields     []string
	Stack      []runtime.Frame
	AfterWrite func()
}

func (e *Entry) log() {
	for _, core := range e.Cores {
		if core == nil {
			continue
		}
		core.Write(e, e.Params...)
	}

	if e.AfterWrite != nil {
		e.AfterWrite()
	}
	switch e.Level {
	case FATL:
		os.Exit(1)
	case PNIC:
		panic(e.Params)
	default:
	}
}
