package core

import (
	"fmt"
	"sort"
	"strings"
)

type Param map[string]any

type Array struct {
	K string
	V any
}

func parseParams(params ...any) any {
	if len(params) == 0 {
		return nil
	}
	format, ok := params[0].(string)
	if ok && strings.ContainsRune(format, '%') {
		return fmt.Sprintf(format, params[1:]...)
	}
	return params
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
