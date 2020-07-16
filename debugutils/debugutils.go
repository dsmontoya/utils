package debugutils

import (
	"runtime/debug"
	"strings"
)

const (
	stackFuncRuntime    = "runtime/debug.Stack"
	stackFuncDebugutils = "github.com/dsmontoya/utils/debugutils.Stack"
)

func Stack() string {
	return string(debug.Stack())
}

func StackSimple() string {
	s := Stack()
	lines := strings.Split(s, "\n")
	newLines := []string{}
	for i := 0; i < len(lines)-1; i++ {
		line := lines[i]
		if i == 0 {
			newLines = append(newLines, line)
		}
		if i%2 == 0 {
			continue
		}
		if strings.Contains(line, stackFuncRuntime) || strings.Contains(line, stackFuncDebugutils) {
			continue
		}
		newLines = append(newLines, line, lines[i+1])
	}
	return strings.Join(newLines, "\n")
}
