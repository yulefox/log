package core

import "encoding/json"

type Level int

const (
	DEBU Level = iota // default level
	INFO
	WARN
	ERRO
	PNIC
	FATL
)

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
