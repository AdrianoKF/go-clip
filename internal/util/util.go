package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func InitializeLogging(development bool) {
	var cfg zap.Config
	if development {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	} else {
		cfg = zap.NewProductionConfig()
	}

	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	Logger = l.Sugar()
}
