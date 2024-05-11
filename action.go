package log

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
}
