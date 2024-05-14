package log

import (
	"fmt"
	"runtime"
	"strings"
)

func TrimPath(file string) string {
	idx := strings.LastIndexByte(file, '/')
	if idx == -1 {
		return file
	}

	return file[idx+1:]
}

func GetStack(skip int, depth int) (stack []string) {
	pc := make([]uintptr, depth)
	n := runtime.Callers(skip, pc)
	if n == 0 {
		return
	}

	frames := runtime.CallersFrames(pc[:n])
	for {
		frame, more := frames.Next()
		caller := fmt.Sprintf("%v %v:%v",
			TrimPath(frame.Function),
			frame.File,
			frame.Line,
		)
		stack = append(stack, caller)
		if !more {
			break
		}
	}

	return
}

func GetStacks(skip int) []string {
	stacks := strings.Split(string(debug.Stack()), "\n")
	if len(stacks) > skip {
		stacks = stacks[skip:]
	}

	return stacks
}
