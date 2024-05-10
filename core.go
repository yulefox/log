package log

import (
	"fmt"
	"io"
)

type Encoder interface {
	Encode(ac *Action, params []any) *Buffer
}

type Core struct {
	allWriter  io.Writer
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
	if c.allWriter != nil {
		fmt.Fprint(c.allWriter, buffer.String(), "\n")
	}
	switch ac.Level {
	case ERRO, PNIC, FATL:
		fmt.Fprint(c.errWriter, buffer.String(), "\n")
	default:
		if c.infoWriter != nil {
			fmt.Fprint(c.infoWriter, buffer.String(), "\n")
		}
	}
}
