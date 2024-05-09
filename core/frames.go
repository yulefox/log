package core

import (
	"fmt"
	"runtime"
)

func GetCallerFrames(skip, dep int) (stack []string) {
	pc := make([]uintptr, dep)
	n := runtime.Callers(skip, pc)

	if n == 0 {
		return
	}

	pc = pc[:n]

	frames := runtime.CallersFrames(pc)

	for {
		frame, more := frames.Next()
		stack = append(stack, fmt.Sprintf("%v:%v %v",
			TrimPath(frame.File),
			frame.Line,
			frame.Function,
		))

		if !more {
			break
		}
	}

	return
}
