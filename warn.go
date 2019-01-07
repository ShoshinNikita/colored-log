package log

import (
	"fmt"
)

// Warn prints warning
// Output pattern: (?time) [WARN] warning
func (l Logger) Warn(v ...interface{}) {
	text := ""
	text = l.getTime()

	text += l.getWarnMsg()
	l.printText(text + fmt.Sprint(v...))
}

// Warnf prints warning
// Output pattern: (?time) [WARN] warning
func (l Logger) Warnf(format string, v ...interface{}) {
	text := ""
	text = l.getTime()

	text += l.getWarnMsg()
	l.printText(text + fmt.Sprintf(format, v...))
}

// Warnln prints warning
// Output pattern: (?time) [WARN] warning
func (l Logger) Warnln(v ...interface{}) {
	text := ""
	text = l.getTime()

	text += l.getWarnMsg()
	l.printText(text + fmt.Sprintln(v...))
}
