package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapErrorLogger(dirName string, level string) (*zap.Logger, error) {
	zapLevel, ok := DictLogLevel[level]
	if !ok {
		zapLevel = zapcore.InfoLevel
	}
	writer, err := timeDivisionWriter(dirName + "/error")
	if err != nil {
		return nil, err
	}
	sink := zapcore.AddSync(writer)
	writeSyncer := zapcore.NewMultiWriteSyncer(sink)
	encoderConfig := GetEncoderConfig()
	encoderConfig.CallerKey = "caller"
	encoderConfig.TimeKey = "ts"
	encoderConfig.LevelKey = "level"
	encoder := zapcore.NewConsoleEncoder(encoderConfig)     //获取编码器,NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	core := zapcore.NewCore(encoder, writeSyncer, zapLevel) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	return zap.New(core, zap.AddCaller()), nil
}

func NewZapAccessLogger(dirName string) (*zap.Logger, error) {
	writer, err := timeDivisionWriter(dirName + "/access")
	if err != nil {
		return nil, err
	}
	sink := zapcore.AddSync(writer)
	writeSyncer := zapcore.NewMultiWriteSyncer(sink)
	encoderConfig := GetEncoderConfig()                               //指定时间格式
	encoder := zapcore.NewConsoleEncoder(encoderConfig)               //获取编码器,NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	return zap.New(core), nil
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
