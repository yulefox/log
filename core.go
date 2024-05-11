package log

import (
	"io"
	"unsafe"
)

type Encoder interface {
	Encode(ac *Action, params []any) string
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

	str := e.Encode(ac, params) + "\n"
	buff := *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{str, len(str)},
	))
	if c.allWriter != nil {
		c.allWriter.Write(buff)
	}
	switch ac.Level {
	case ERRO, PNIC, FATL:
		if c.errWriter != nil {
			c.errWriter.Write(buff)
		}
	default:
		if c.infoWriter != nil {
			c.infoWriter.Write(buff)
		}
	}
}
