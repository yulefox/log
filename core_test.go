package log

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCoreWrite(t *testing.T) {
	var buf bytes.Buffer
	core := &Core{
		allWriter: &buf,
		encoder:   &TestEncoder{},
	}

	core.Write(&Entry{Level: INFO, Date: "2020-01-01"}, "info message")

	expected := "info message\n"
	if buf.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, buf.String())
	}
}

func TestCoreWriteWithNilEncoder(t *testing.T) {
	var buf bytes.Buffer
	core := &Core{
		allWriter:  &buf,
		infoWriter: &buf,
		errWriter:  &buf,
		encoder:    nil,
	}

	core.Write(&Entry{Level: INFO, Date: "2020-01-01"})

	expected := ""
	if buf.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, buf.String())
	}
}

func TestCoreWriteWithNilEntry(t *testing.T) {
	var buf bytes.Buffer
	core := &Core{
		allWriter:  &buf,
		infoWriter: &buf,
		errWriter:  &buf,
		encoder:    &TestEncoder{},
	}

	core.Write(nil)

	expected := ""
	if buf.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, buf.String())
	}
}

type TestEncoder struct{}

func (e *TestEncoder) Encode(ac *Entry, params []any) string {
	w := bufferPool.Get().(*Buffer)
	defer w.close()
	if _, err := fmt.Fprint(w, params...); err != nil {
		return ""
	}
	return w.String()
}
