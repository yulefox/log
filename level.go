package log

import (
	"encoding/json"
	"strings"
)

type Level uint32

const (
	DEBU Level = iota // default level
	INFO
	WARN
	ERRO
	PNIC
	FATL
)

func StringToLevel(logLevel string) Level {

	switch strings.ToLower(logLevel) {
	case "panic", "pani":
		return PNIC
	case "fatal", "fata":
		return FATL
	case "error":
		return ERRO
	case "warning", "warn":
		return WARN
	case "info":
		return INFO
	case "debug", "debu":
		return DEBU
	default:
		return INFO
	}
}

func (l Level) String() string {
	switch l {
	case DEBU:
		return "DEBU"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERRO:
		return "ERRO"
	case PNIC:
		return "PNIC"
	case FATL:
		return "FATL"
	default:
		return "UNKN"
	}
}

func (l Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}
