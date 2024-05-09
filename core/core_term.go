package core

import (
	"fmt"
	"os"
	"strings"
)

type TermEncoder struct{}

func NewTermCore() *Core {
	return &Core{
		infoWriter: os.Stdout,
		errWriter:  os.Stderr,
		encoder:    new(TermEncoder),
	}
}

func (e *TermEncoder) Encode(ac *Action, params []any) *Buffer {
	if ac == nil {
		return nil
	}

	w := bufferPool.Get().(*Buffer)
	shader := shaderByLv(ac.Level)
	if shader != nil {
		w.WriteString(shader.do(ac.Level.String()))
	} else {
		w.WriteString(ac.Level.String())
	}

	if ac.Date != "" {
		w.WriteString(" " + ac.Date)
	}

	if ac.AddCaller && ac.Caller != "" {
		w.WriteString(" " + ac.Caller)
	}

	if ac.Name != "" {
		w.WriteString(" [" + ac.Name + "]")
	}

	if params != nil {
		w.WriteString(" ")
		format, ok := params[0].(string)
		if ok && strings.ContainsRune(format, '%') {
			if _, err := fmt.Fprintf(w, format, params[1:]...); err != nil {
				return w
			}
		} else {
			if _, err := fmt.Fprint(w, params...); err != nil {
				return w
			}
		}
	}

	for i, layer := range ac.Stack {
		if _, err := fmt.Fprintf(w, "\n\033[31m%2v %v\033[0m", i+1, layer); err != nil {
			return w
		}
	}

	return w
}
