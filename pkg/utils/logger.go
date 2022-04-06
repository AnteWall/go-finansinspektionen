package utils

import (
	"github.com/onsi/ginkgo/reporters/stenographer/support/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.SugaredLogger {
	if IsDevEnvironment() {
		aa := zap.NewDevelopmentEncoderConfig()
		aa.EncodeLevel = zapcore.CapitalColorLevelEncoder
		devLogger := zap.New(zapcore.NewCore(
			zapcore.NewConsoleEncoder(aa),
			zapcore.AddSync(colorable.NewColorableStdout()),
			zapcore.DebugLevel,
		))
		return devLogger.Sugar()
	}
	log, _ := zap.NewProduction()
	return log.Sugar()
}
