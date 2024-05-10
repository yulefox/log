package log

import (
	"fmt"
	"sort"
	"strings"
)

type Action struct {
	*Options
	Level      Level
	Date       string
	Caller     string
	Stack      []string
	AfterWrite func()
}

type Param map[string]any

type Array struct {
	K string
	V any
}

func (p Param) String() string {
	array := make([]Array, 0, len(p))

	for k, v := range p {
		array = append(array, Array{K: k, V: v})
	}

	sort.Slice(array, func(i, j int) bool {
		return array[i].K < array[j].K
	})

	var result strings.Builder

	for _, v := range array {
		result.WriteString(fmt.Sprintf("%v:%+v ", v.K, v.V))
	}

	return strings.TrimSuffix(result.String(), " ")
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
