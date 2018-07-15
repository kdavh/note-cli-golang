package nreport

import (
	"fmt"
)

type Reporter struct {
	level uint
}

func (l *Reporter) Debugf(str string, args ...interface{}) {
	if l.level >= DEBUG {
		fmt.Printf(str, args...)
	}
}

func (l *Reporter) Infof(str string, args ...interface{}) {
	if l.level >= INFO {
		fmt.Printf(str, args...)
	}
}

func (l *Reporter) Errorf(str string, args ...interface{}) {
	if l.level >= ERROR {
		fmt.Printf(str, args...)
	}
}

func (l *Reporter) Error(str string) {
	if l.level >= ERROR {
		fmt.Println(str)
	}
}

func (l *Reporter) Reportf(str string, args ...interface{}) {
	fmt.Printf(str, args...)
}

func (l *Reporter)Prompt() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func New(level uint) *Reporter {
	return &Reporter{
		level: level,
	}
}

const DEBUG = 8
const INFO = 5
const ERROR = 2
