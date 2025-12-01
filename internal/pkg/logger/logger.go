package logger

import "go.uber.org/zap"

var Log *zap.Logger

func InitZap(serviceName string) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}

	config.EncoderConfig.TimeKey = "time"

	Log, _ = config.Build(zap.Fields(zap.String("service", serviceName)))
}
