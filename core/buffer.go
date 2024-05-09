package core

import (
	"strings"
	"sync"
)

type Buffer struct {
	strings.Builder
}

var (
	bufferPool = sync.Pool{
		New: func() interface{} {
			return new(Buffer)
		},
	}
)

func (b *Buffer) close() {
	b.Reset()
	bufferPool.Put(b)
}
