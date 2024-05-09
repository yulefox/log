package core

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
