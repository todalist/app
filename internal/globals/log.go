package globals

import (
	"go.uber.org/zap"
)

var (
	LOG *zap.Logger
)

func InitLogging() (_ *zap.Logger, err error) {
	isProd := IsProd()
	if IsProd() {
		LOG, err = zap.NewProduction()
	} else {
		LOG, err = zap.NewDevelopment()
	}
	LOG.Info("new logging", zap.Bool("production", isProd))
	return LOG, err
}
