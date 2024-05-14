package log

import (
	"runtime/debug"
	"strings"
)

func TrimPath(file string) string {
	idx := strings.LastIndexByte(file, '/')
	if idx == -1 {
		return file
	}

	idx = strings.LastIndexByte(file[:idx], '/')
	if idx == -1 {
		return file
	}

	return file[idx+1:]
}

func GetStacks(skip int) []string {
	stacks := strings.Split(string(debug.Stack()), "\n")
	if len(stacks) > skip {
		stacks = stacks[skip:]
	}

	return stacks
}
