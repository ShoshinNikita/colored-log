// Package log provides functions for pretty print
//
// Patterns of functions print:
// * Print(), Printf(), Println():
//   (?time) msg
// * Info(), Infof(), Infoln():
//   (?time) [INFO] msg
// * Warn(), Warnf(), Warnln():
//   (?time) [WARN] msg
// * Error(), Errorf(), Errorln():
//   (?time) [ERR] (?file:line) error
// * Fatal(), Fatalf(), Fatalln():
//   (?time) [FATAL] (?file:line) error
//
// Time pattern: MM.dd.yyyy hh:mm:ss (01.30.2018 05:5:59)
//
package log

import (
	"fmt"

	"github.com/fatih/color"
)

const (
	timeLayout = "01.02.2006 15:04:05"
)

var (
	// For time
	timePrintf = color.New(color.FgHiGreen).SprintfFunc()

	// For [INFO]
	infoPrint = color.New(color.FgCyan).SprintFunc()

	// For [WARN]
	warnPrint = color.New(color.FgYellow).SprintFunc()

	// For [ERR]
	errorPrint = color.New(color.FgRed).SprintFunc()

	// For [FATAL]
	fatalPrint = color.New(color.BgRed).SprintFunc()
)

// init inits globalLogger with NewLogger()
func init() {
	globalLogger = NewLogger()

	globalLogger.PrintTime = false
	globalLogger.PrintColor = true
	globalLogger.PrintErrorLine = true
}

type textStruct struct {
	text string
	ch   chan struct{}
}

func newText(text string) textStruct {
	return textStruct{text: text, ch: make(chan struct{})}
}

func (t *textStruct) done() {
	close(t.ch)
}

type Logger struct {
	PrintTime      bool
	PrintColor     bool
	PrintErrorLine bool

	printChan chan textStruct
}

// NewLogger creates *Logger and run goroutine (Logger.printer())
func NewLogger() *Logger {
	l := new(Logger)
	l.printChan = make(chan textStruct)
	go l.printer()
	return l
}

func (l *Logger) printer() {
	for text := range l.printChan {
		fmt.Fprint(color.Output, text.text)
		text.done()
	}
}

func (l *Logger) printText(text string) {
	t := newText(text)
	l.printChan <- t
	<-t.ch
}

var globalLogger *Logger

// PrintTime sets globalLogger.PrintTime
// Time isn't printed by default
func PrintTime(b bool) {
	globalLogger.PrintTime = b
}

// ShowTime sets printTime
// Time isn't printed by default
//
// It was left for backwards compatibility
var ShowTime = PrintTime

// PrintColor sets printColor
// printColor is true by default
func PrintColor(b bool) {
	globalLogger.PrintColor = b
}

// PrintErrorLine sets PrintErrorLine
// If PrintErrorLine is true, log.Error(), log.Errorf(), log.Errorln() will print file and line,
// where functions were called.
// PrintErrorLine is true by default
func PrintErrorLine(b bool) {
	globalLogger.PrintErrorLine = b
}

/* Print */

func Print(v ...interface{}) {
	globalLogger.Print(v...)
}

func Printf(format string, v ...interface{}) {
	globalLogger.Printf(format, v...)
}

func Println(v ...interface{}) {
	globalLogger.Println(v...)
}

/* Info */

func Info(v ...interface{}) {
	globalLogger.Info(v...)
}

func Infof(format string, v ...interface{}) {
	globalLogger.Infof(format, v...)
}

func Infoln(v ...interface{}) {
	globalLogger.Infoln(v...)
}

/* Warn */

func Warn(v ...interface{}) {
	globalLogger.Warn(v...)
}

func Warnf(format string, v ...interface{}) {
	globalLogger.Warnf(format, v...)
}

func Warnln(v ...interface{}) {
	globalLogger.Warnln(v...)
}

/* Error */

// Error prints error
// Output pattern: (?time) [ERR] (?file:line) error
func Error(v ...interface{}) {
	globalLogger.Error(v...)
}

// Errorf prints error
// Output pattern: (?time) [ERR] (?file:line) error
func Errorf(format string, v ...interface{}) {
	globalLogger.Errorf(format, v...)
}

// Errorln prints error
// Output pattern: (?time) [ERR] (?file:line) error
func Errorln(v ...interface{}) {
	globalLogger.Errorln(v...)
}

/* Fatal */

func Fatal(v ...interface{}) {
	globalLogger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	globalLogger.Fatalf(format, v...)
}

func Fatalln(v ...interface{}) {
	globalLogger.Fatalln(v...)
}
