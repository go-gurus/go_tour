// zap_facade/zap_facade.go
package zap_facade

import (
	"go.uber.org/zap"
)

type ZapStruct struct {
	logger zap.Logger
}

func (s ZapStruct) Info(msg string) {
	defer s.logger.Sync()
	s.logger.Info(msg + " (zap)")
}

func (s ZapStruct) Error(msg string) {
	defer s.logger.Sync()
	s.logger.Error(msg + " (zap)")
}

func NewZapStruct() ZapStruct {
	logger, _ := zap.NewProduction()
	result := ZapStruct{
		*logger,
	}
	return result
}
