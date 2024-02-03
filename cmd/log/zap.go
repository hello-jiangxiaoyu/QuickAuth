package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapErrorLogger(dirName string) *zap.Logger {
	writer := getLumberjackWriter(dirName + "/error")
	sink := zapcore.AddSync(writer)
	writeSyncer := zapcore.NewMultiWriteSyncer(sink)
	encoderConfig := GetEncoderConfig()
	encoderConfig.CallerKey = "caller"
	encoderConfig.TimeKey = "ts"
	encoderConfig.LevelKey = "level"
	encoder := zapcore.NewConsoleEncoder(encoderConfig)              // 获取编码器
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel) // 日志级别
	return zap.New(core, zap.AddCaller())
}

func NewZapAccessLogger(dirName string) *zap.Logger {
	writer := getLumberjackWriter(dirName + "/access")
	sink := zapcore.AddSync(writer)
	writeSyncer := zapcore.NewMultiWriteSyncer(sink)
	encoderConfig := GetEncoderConfig()
	encoder := zapcore.NewConsoleEncoder(encoderConfig)               // 获取编码器
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel) // 日志级别
	return zap.New(core)
}

func GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		NameKey:        "logger",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
