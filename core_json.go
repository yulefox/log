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

func (e *JsonEncoder) Encode(entry *Entry, params []any) string {
	if entry == nil {
		return ""
	}

	buf, err := json.Marshal(&structureLog{
		Date:   entry.Date,
		Level:  entry.Level,
		Caller: entry.Caller,
		Fields: entry.Fields,
		Params: params,
	})

	if err != nil {
		return ""
	}

	return unsafe.String(unsafe.SliceData(buf), len(buf))
}
