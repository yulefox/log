package core

import (
	"fmt"
	"io"
)

type Encoder interface {
	Encode(ac *Action, params []any) *Buffer
}

type Core struct {
	infoWriter io.Writer
	errWriter  io.Writer
	encoder    Encoder
}

func (c *Core) Write(ac *Action, params ...any) {
	e := c.encoder
	if e == nil || ac == nil {
		return
	}

	buffer := e.Encode(ac, params)
	if buffer == nil {
		return
	}

	defer buffer.close()
	switch ac.Level {
	case ERRO, PNIC, FATL:
		fmt.Fprint(c.errWriter, buffer.String(), "\n")
	default:
		fmt.Fprint(c.infoWriter, buffer.String(), "\n")
	}
}
