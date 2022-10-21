// logrus_facade/logrus_facade.go
package logrus_facade

import (
	"github.com/sirupsen/logrus"
)

type LogrusStruct struct{}

func (logger LogrusStruct) Info(msg string) {
	logrus.Info(msg + " (logrus)")
}

func (logger LogrusStruct) Error(msg string) {
	logrus.Error(msg + " (logrus)")
}
