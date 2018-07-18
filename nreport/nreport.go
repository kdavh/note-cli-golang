package nreport

import (
	"fmt"
)

type Reporter struct {
	Level uint
}

func (l *Reporter) Debugf(str string, args ...interface{}) {
	if l.Level >= DEBUG {
		fmt.Printf(str, args...)
	}
}

func (l *Reporter) Infof(str string, args ...interface{}) {
	if l.Level >= INFO {
		fmt.Printf(str, args...)
	}
}

func (l *Reporter) Errorf(str string, args ...interface{}) {
	if l.Level >= ERROR {
		fmt.Printf(str, args...)
	}
}

func (l *Reporter) Error(str string) {
	if l.Level >= ERROR {
		fmt.Println(str)
	}
}

func (l *Reporter) Reportf(str string, args ...interface{}) {
	fmt.Printf(str, args...)
}

func (l *Reporter) Prompt() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func New(level uint) *Reporter {
	return &Reporter{
		Level: level,
	}
}

const DEBUG = 8
const INFO = 5
const ERROR = 2
