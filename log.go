// Package log provides functions for pretty print
//
// Patterns of functions print:
// * Print(), Printf(), Println():
//   (?time) msg
// * Info(), Infof(), Infoln():
//   (?time) [INFO] msg
// * Warn(), Warnf(), Warnln():
//   (?time) [WARN] warning
// * Error(), Errorf(), Errorln():
//   (?time) [ERR] (?file:line) error
// * Fatal(), Fatalf(), Fatalln():
//   (?time) [FATAL] (?file:line) error
//
// Time pattern: MM.dd.yyyy hh:mm:ss (01.30.2018 05:5:59)
//
package log

import (
	"io"
	"log"

	"github.com/fatih/color"
)

const (
	DefaultTimeLayout = "01.02.2006 15:04:05"
)

type Logger struct {
	printTime      bool
	printColor     bool
	printErrorLine bool

	global bool

	output     io.Writer
	timeLayout string
}

// NewLogger creates *Logger and run goroutine (Logger.printer())
func NewLogger() *Logger {
	l := new(Logger)
	l.output = color.Output
	l.timeLayout = DefaultTimeLayout
	return l
}

func (l *Logger) printText(text string) {
	log.Print(text)
}

// PrintTime sets Logger.printTime to b
func (l *Logger) PrintTime(b bool) {
	l.printTime = b
}

// PrintColor sets Logger.printColor to b
func (l *Logger) PrintColor(b bool) {
	l.printColor = b
}

// PrintErrorLine sets Logger.printErrorLine to b
func (l *Logger) PrintErrorLine(b bool) {
	l.printErrorLine = b
}

// ChangeOutput changes Logger.output writer.
// Default Logger.output is github.com/fatih/color.Output
func (l *Logger) ChangeOutput(w io.Writer) {
	// l.output = w
	log.SetOutput(w)
}

// ChangeTimeLayout changes Logger.timeLayout
// Default Logger.timeLayout is DefaultTimeLayout
func (l *Logger) ChangeTimeLayout(layout string) {
	l.timeLayout = layout
}
