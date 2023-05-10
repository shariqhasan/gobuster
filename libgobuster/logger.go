package libgobuster

import (
	"log"
	"os"

	"github.com/fatih/color"
)

type Logger struct {
	log      *log.Logger
	errorLog *log.Logger
	debug    bool
}

func NewLogger(debug bool) Logger {
	return Logger{
		log:      log.New(os.Stdout, "", log.LstdFlags),
		errorLog: log.New(os.Stderr, color.New(color.FgRed).Sprint("[ERROR] "), log.LstdFlags),
		debug:    debug,
	}
}

func (l Logger) Debug(v ...any) {
	if !l.debug {
		return
	}
	l.log.Print(v...)
}

func (l Logger) Debugf(format string, v ...any) {
	if !l.debug {
		return
	}
	l.log.Printf(format, v...)
}

func (l Logger) Print(v ...any) {
	l.log.Print(v...)
}

func (l Logger) Printf(format string, v ...any) {
	l.log.Printf(format, v...)
}

func (l Logger) Println(v ...any) {
	l.log.Println(v...)
}

func (l Logger) Error(v ...any) {
	l.errorLog.Print(v...)
}

func (l Logger) Errorf(format string, v ...any) {
	l.errorLog.Printf(format, v...)
}

func (l Logger) Fatal(v ...any) {
	l.errorLog.Fatal(v...)
}

func (l Logger) Fatalf(format string, v ...any) {
	l.errorLog.Fatalf(format, v...)
}

func (l Logger) Fatalln(v ...any) {
	l.errorLog.Fatalln(v...)
}
