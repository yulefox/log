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

func (e *TermEncoder) Encode(ac *Entry, params []any) string {
	if ac == nil {
		return ""
	}

	w := bufferPool.Get().(*Buffer)
	defer w.close()

	if ac.Date != "" {
		w.WriteString(ac.Date + " ")
	}
	shader := shaderByLv(ac.Level)
	if shader != nil {
		w.WriteString(shader.do(ac.Level.String()))
	} else {
		w.WriteString(ac.Level.String())
	}
	if len(ac.Fields) > 0 {
		w.WriteString(" [" + strings.Join(ac.Fields, " ") + "]")
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

	if ac.AddCaller && ac.Caller != "" {
		w.WriteString(" " + ac.Caller)
	}
	for i, layer := range ac.Stack {
		if _, err := fmt.Fprintf(w, "\n\033[31m%2v %v\033[0m", i+1, layer); err != nil {
			return ""
		}
	}

	return w.String()
}
