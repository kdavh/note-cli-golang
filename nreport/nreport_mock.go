package nreport

import (
	"fmt"
)

type ReporterMock struct {
	DebugCalls []string
	InfoCalls  []string
	ErrorCalls []string
}

func (l *ReporterMock) Debugf(str string, args ...interface{}) {
	l.DebugCalls = append(l.DebugCalls, fmt.Sprintf(str, args...))
}

func (l *ReporterMock) Infof(str string, args ...interface{}) {
	l.InfoCalls = append(l.InfoCalls, fmt.Sprintf(str, args...))
}

func (l *ReporterMock) Errorf(str string, args ...interface{}) {
	l.ErrorCalls = append(l.ErrorCalls, fmt.Sprintf(str, args...))
}

func (l *ReporterMock) Error(str string) {
	l.ErrorCalls = append(l.ErrorCalls, fmt.Sprintln(str))
}

func (l *ReporterMock) Reset() {
	l.DebugCalls = []string{}
	l.InfoCalls = []string{}
	l.ErrorCalls = []string{}
}

func NewMock() *ReporterMock {
	return &ReporterMock{
		DebugCalls: []string{},
		InfoCalls:  []string{},
		ErrorCalls: []string{},
	}
}
