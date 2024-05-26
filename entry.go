package log

import (
	"os"
	"runtime"
)

type Entry struct {
	*Options
	Level      Level
	Date       string
	Caller     string
	Fields     []string
	Stack      []runtime.Frame
	AfterWrite func()
}

func (e *Entry) log(level Level, params ...any) {
	if level < e.Options.Level {
		return
	}
	e.Level = level

	if e.TimeFormat != "" {
		e.Date = e.Now().Format(e.TimeFormat)
	}

	if e.AddCaller || e.Level >= ERRO {
		stack := getStack(e.Skip, 10)
		if len(stack) > 0 {
			if e.AddCaller {
				e.Caller = e.FormatFrame(stack[0])
			}
			switch e.Level {
			case ERRO, FATL, PNIC:
				e.Stack = stack
			default:
			}
		}
	}

	for _, core := range e.Cores {
		if core == nil {
			continue
		}
		core.Write(e, params...)
	}

	if e.AfterWrite != nil {
		e.AfterWrite()
	}
	switch e.Level {
	case FATL:
		os.Exit(1)
	case PNIC:
		panic(params)
	default:
	}
}
