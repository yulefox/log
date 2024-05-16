package log

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type logEntry struct {
	level  Level
	params []any
	caller string
	stack  []string
}

type Entry struct {
	*Options
	Level      Level
	Date       string
	Caller     string
	Fields     []string
	Stack      []string
	channel    chan *logEntry
	AfterWrite func()
}

func toString(frame runtime.Frame) string {
	return fmt.Sprintf("%v %v:%v", frame.Function, frame.File, frame.Line)
}

func toStrings(frames []runtime.Frame) []string {
	stack := make([]string, 0, len(frames))
	for _, frame := range frames {
		stack = append(stack, toString(frame))
	}
	return stack
}

func (e *Entry) log(level Level, params ...any) {
	if level < e.Options.Level {
		return
	}
	en := &logEntry{level: level, params: params}
	if e.AddCaller || e.Level >= ERRO {
		depth := 1
		if e.Level >= ERRO {
			depth = 10
		}
		stack := GetStack(e.Skip, depth)
		if len(stack) > 0 {
			if e.AddCaller {
				en.caller = toString(stack[0])
			}

			switch e.Level {
			case ERRO, FATL, PNIC:
				en.stack = toStrings(stack)
			default:
			}
		}
	}
	select {
	case e.channel <- en:
	case <-time.After(3 * time.Second):
	}
}

func (e *Entry) _log(en *logEntry) {
	e.Level = en.level
	e.Caller = en.caller
	e.Stack = en.stack

	if e.TimeFormat != "" {
		e.Date = e.Now().Format(e.TimeFormat) //+ " real: " + time.Now().Format(e.TimeFormat)
	}

	for _, core := range e.Cores {
		if core == nil {
			continue
		}
		core.Write(e, en.params...)
	}

	if e.AfterWrite != nil {
		e.AfterWrite()
	}
	switch e.Level {
	case FATL:
		os.Exit(1)
	case PNIC:
		panic(en.params)
	default:
	}
}
