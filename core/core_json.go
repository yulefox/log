package core

import (
	"encoding/json"
	"os"
)

type JsonEncoder struct {
}

func NewJsonCore() *Core {
	return &Core{
		infoWriter: os.Stdout,
		errWriter:  os.Stderr,
		encoder:    new(JsonEncoder),
	}
}

func (e *JsonEncoder) Encode(ac *Action, params []any) *Buffer {
	if ac == nil {
		return nil
	}

	s := action2Structure(ac)
	s.Params = params

	data, err := json.Marshal(s)

	if err != nil {
		return nil
	}

	cache := bufferPool.Get().(*Buffer)
	if _, err := cache.Write(data); err != nil {
		return nil
	}

	return cache
}
