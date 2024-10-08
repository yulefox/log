package log

import (
	"fmt"
	"runtime"
)

type Entry struct {
	*Options
	Level      Level
	Params     []any
	Date       string
	Caller     string
	Fields     []string
	Stack      []string
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
}
