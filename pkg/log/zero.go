package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewZeroErrorLog(dirName string) (*zerolog.Logger, error) {
	writer, err := timeDivisionWriter(dirName + "/error")
	if err != nil {
		return nil, err
	}

	consoleWriter := zerolog.ConsoleWriter{
		Out:     writer,
		NoColor: true,
	}
	logger := log.Output(consoleWriter).With().Logger()
	return &logger, nil
}

func NewZeroAccessLog(dirName string) (*zerolog.Logger, error) {
	writer, err := timeDivisionWriter(dirName + "/access")
	if err != nil {
		return nil, err
	}

	consoleWriter := zerolog.ConsoleWriter{
		Out:             writer,
		NoColor:         true,
		FormatTimestamp: func(i interface{}) string { return "" },
		FormatLevel:     func(i interface{}) string { return "" },
	}
	logger := log.Output(consoleWriter).With().Logger()
	return &logger, nil
}
