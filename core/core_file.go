package core

import (
	"fmt"
	"os"
	"strings"
)

type FileEncoder struct{}

func NewFileCore(name string) *Core {
	_ = os.MkdirAll("./logs", 0755)
	infoFile, err := os.OpenFile(fmt.Sprintf("logs/%s.log", name), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	errFile, err := os.OpenFile(fmt.Sprintf("logs/%s_err.log", name), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	return &Core{
		infoWriter: infoFile,
		errWriter:  errFile,
		encoder:    new(FileEncoder),
	}
}

func (e *FileEncoder) Encode(ac *Action, params []any) *Buffer {
	if ac == nil {
		return nil
	}

	w := bufferPool.Get().(*Buffer)

	if ac.Date != "" {
		w.WriteString(ac.Date + " ")
	}
	w.WriteString(ac.Level.String())
	if ac.AddCaller && ac.Caller != "" {
		w.WriteString(" " + ac.Caller + " ")
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
		if _, err := fmt.Fprintf(w, "\n %2d %v", i+1, layer); err != nil {
			return w
		}
	}

	return w
}
