package log

import (
	"encoding/json"
	"io"
	"unsafe"
)

type JsonEncoder struct {
}

type structureLog struct {
	Level  Level       `json:"level"`
	Date   string      `json:"date,omitempty"`
	Caller string      `json:"caller,omitempty"`
	Params interface{} `json:"params,omitempty"`
	Fields []string    `json:"fields,omitempty"`
	Stack  []string    `json:"stack,omitempty"`
}

func NewJsonCore(writer io.Writer) *Core {
	return &Core{
		allWriter: writer,
		encoder:   new(JsonEncoder),
	}
}

func (e *JsonEncoder) Encode(ac *Entry, params []any) string {
	if ac == nil {
		return ""
	}

	buf, err := json.Marshal(&structureLog{
		Date:   ac.Date,
		Level:  ac.Level,
		Caller: ac.Caller,
		Fields: ac.Fields,
		Stack:  ac.Stack,
		Params: params,
	})

	if err != nil {
		return ""
	}

	return unsafe.String(unsafe.SliceData(buf), len(buf))
}
