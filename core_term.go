package log

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

func (e *TermEncoder) Encode(entry *Entry, params []any) string {
	if entry == nil {
		return ""
	}

	w := bufferPool.Get().(*Buffer)
	defer w.close()

	if entry.Date != "" {
		w.WriteString(entry.Date + " ")
	}
	shader := shaderByLv(entry.Level)
	if shader != nil {
		w.WriteString(shader.do(entry.Level.String()))
	} else {
		w.WriteString(entry.Level.String())
	}
	if entry.AddCaller && entry.Caller != "" {
		w.WriteString(" " + entry.Caller)
	}
	if len(entry.Fields) > 0 {
		w.WriteString(" [" + strings.Join(entry.Fields, " ") + "]")
	}

	if params != nil {
		w.WriteString(" ")
		format, ok := params[0].(string)
		if ok && strings.ContainsRune(format, '%') {
			if _, err := fmt.Fprintf(w, format, params[1:]...); err != nil {
				return ""
			}
		} else {
			if _, err := fmt.Fprint(w, params...); err != nil {
				return ""
			}
		}
	}

	for i, layer := range entry.Stack {
		if _, err := fmt.Fprintf(w, "\n\033[31m%2v %v\033[0m", i+1, layer); err != nil {
			return ""
		}
	}

	return w.String()
}
