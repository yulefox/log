package log

import (
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

func GetStack(skip int, depth int) (stack []runtime.Frame) {
	pc := make([]uintptr, depth)
	n := runtime.Callers(skip, pc)
	if n == 0 {
		return
	}

	frames := runtime.CallersFrames(pc[:n])
	for {
		frame, more := frames.Next()
		stack = append(stack, frame)
		if !more {
			break
		}
	}

	return
}
