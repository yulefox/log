package log

import (
	"io"
	"unsafe"
)

type Encoder interface {
	Encode(entry *Entry, params []any) string
}

type Core struct {
	allWriter  io.Writer
	infoWriter io.Writer
	errWriter  io.Writer
	encoder    Encoder
}

func (c *Core) Write(entry *Entry, params ...any) {
	e := c.encoder
	if e == nil || entry == nil {
		return
	}

	str := e.Encode(entry, params) + "\n"
	buff := *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{str, len(str)},
	))
	if c.allWriter != nil {
		c.allWriter.Write(buff)
	}
	switch entry.Level {
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
