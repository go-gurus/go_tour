// golog_facade/golog_facade.go
package golog_facade

import (
	"github.com/kataras/golog"
)

type GologStruct struct{}

func (s GologStruct) Info(msg string) {
	golog.Info(msg + " (golog)")
}

func (s GologStruct) Error(msg string) {
	golog.Error(msg + " (golog)")
}
