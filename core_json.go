package log

import (
	"encoding/json"
	"os"
)

type JsonEncoder struct {
}

type structureLog struct {
	Name     string      `json:"name,omitempty"`
	Level    Level       `json:"level"`
	Date     string      `json:"date,omitempty"`
	Caller   string      `json:"caller,omitempty"`
	Category string      `json:"category"`
	Reason   string      `json:"reason"`
	Params   interface{} `json:"params,omitempty"`
	Stack    []string    `json:"stack,omitempty"`
}

func action2Structure(ac *Action) *structureLog {
	return &structureLog{
		Name:   ac.Name,
		Date:   ac.Date,
		Level:  ac.Level,
		Caller: ac.Caller,
		Stack:  ac.Stack,
	}
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
