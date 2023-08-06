package log

import (
	rotate "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

var DictLogLevel = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
}

func timeDivisionWriter(path string) (io.Writer, error) {
	MaxAge := 3
	hook, err := rotate.New(
		path+"_%Y-%m-%d.log",
		rotate.WithMaxAge(time.Duration(int64(24*time.Hour)*int64(MaxAge))),
		rotate.WithRotationTime(time.Minute),
		rotate.WithLinkName(path), // 最新日志软链接
	)

	if err != nil {
		return nil, err
	}
	return hook, nil
}
