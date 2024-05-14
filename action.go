package log

import "os"

type Action struct {
	*Options
	Tag        string
	Level      Level
	Date       string
	Caller     string
	Stack      []string
	AfterWrite func()
}

func (a *Action) Write(params ...any) {
	for _, core := range a.Cores {
		if core == nil {
			continue
		}
		core.Write(a, params...)
	}

	if a.AfterWrite != nil {
		a.AfterWrite()
	}
	switch a.Level {
	case FATL:
		os.Exit(1)
	case PNIC:
		panic(params)
	default:
	}
}
