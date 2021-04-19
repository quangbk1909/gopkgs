package stacks

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	maxDepth = 30

	frameFormat  = "\n%s:%d"
	callerFormat = "%s:%d"
)

func DebugInformation(skip int) (stack string, caller string) {
	return getDebugInformation(skip)
}

func Caller(skip int) string {
	_, c := getDebugInformation(skip)
	return c
}

func Stack(skip int) string {
	s, _ := getDebugInformation(skip)
	return s
}

func getDebugInformation(skip int) (stack string, caller string) {
	const callerSkip = 3
	var (
		s, c      strings.Builder
		gotCaller = false
	)
	pc := make([]uintptr, maxDepth)
	n := runtime.Callers(skip+callerSkip, pc)
	if n < len(pc) {
		pc = pc[:n]
	}
	frames := runtime.CallersFrames(pc)
	for frame, more := frames.Next(); frame.PC != 0; frame, more = frames.Next() {
		if !gotCaller {
			c.WriteString(fmt.Sprintf(callerFormat, frame.File, frame.Line))
			gotCaller = true
		}
		s.WriteString(fmt.Sprintf(frameFormat, frame.File, frame.Line))
		if !more {
			break
		}
	}
	return s.String(), c.String()
}
