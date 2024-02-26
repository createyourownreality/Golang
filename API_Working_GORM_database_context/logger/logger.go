package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig

	config.EncoderConfig.StacktraceKey = " "

	log, err = config.Build() // this returns the logger

	if err != nil {
		panic(err)
	}
}

// the three dots in the zap.Field is like that arguement can have any number of variables in it
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}
