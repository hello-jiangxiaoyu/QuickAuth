package initial

import (
	"QuickAuth/conf"
	rotate "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

func timeDivisionWriter(path string) io.Writer {
	MaxAge := 7
	hook, err := rotate.New(
		path+"_%Y-%m-%d.log",
		rotate.WithMaxAge(time.Duration(int64(24*time.Hour)*int64(MaxAge))),
		rotate.WithRotationTime(time.Minute),
		rotate.WithLinkName(path),
	)

	if err != nil {
		panic(err)
	}
	return hook
}

func NewZapErrorLogger(c *conf.SysConfig) *zap.Logger {
	writer := timeDivisionWriter(c.Log.Dir + "/error")
	sink := zapcore.AddSync(writer)
	writeSyncer := zapcore.NewMultiWriteSyncer(sink)
	encoderConfig := GetEncoderConfig()
	encoderConfig.CallerKey = "caller"
	encoderConfig.TimeKey = "ts"
	encoderConfig.LevelKey = "level"
	encoder := zapcore.NewConsoleEncoder(encoderConfig)               //获取编码器,NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	return zap.New(core, zap.AddCaller())
}

func NewZapAccessLogger(c *conf.SysConfig) *zap.Logger {
	writer := timeDivisionWriter(c.Log.Dir + "/access")
	sink := zapcore.AddSync(writer)
	writeSyncer := zapcore.NewMultiWriteSyncer(sink)
	encoderConfig := GetEncoderConfig()                               //指定时间格式
	encoder := zapcore.NewConsoleEncoder(encoderConfig)               //获取编码器,NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
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
