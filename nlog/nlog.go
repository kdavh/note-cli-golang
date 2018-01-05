package nlog

import (
	"fmt"
)

type Logger struct {
	level uint
}

func (l *Logger) Debugf(str string, args ...interface{}) bool {
	if l.level >= DEBUG {
		fmt.Printf(str, args...)
	}
	return true
}

func (l *Logger) Errorf(str string, args ...interface{}) bool {
	if l.level >= ERROR {
		fmt.Printf(str, args...)
	}
	return true
}

func (l *Logger) Error(str string) bool {
	if l.level >= ERROR {
		fmt.Println(str)
	}
	return true
}

func New(level uint) *Logger {
	return &Logger{
		level: level,
	}
}

const DEBUG = 8
const INFO = 5
const ERROR = 2
