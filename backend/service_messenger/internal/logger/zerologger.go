package logger

import (
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/lumberjack.v3"
)

type ZeroLogger struct {
	logger zerolog.Logger
}

func NewZeroLogger(logFolderPath string, logFileName string) (*ZeroLogger, error) {
	err := prepareLogFolder(logFolderPath)
	if err != nil {
		return nil, err
	}

	fileHandler, err := lumberjack.New(
		lumberjack.WithFileName(filepath.Join(logFolderPath, logFileName)),
		lumberjack.WithMaxBytes(1*lumberjack.MB),
		lumberjack.WithMaxBackups(3),
		lumberjack.WithMaxDays(28),
		lumberjack.WithCompress(),
	)
	if err != nil {
		return nil, err
	}

	logger := zerolog.New(zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}, fileHandler)).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Int("pid", os.Getpid()).
		Logger()

	return &ZeroLogger{logger}, nil
}

func (l *ZeroLogger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

func (l *ZeroLogger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *ZeroLogger) Warning(msg string) {
	l.logger.Warn().Msg(msg)
}

func (l *ZeroLogger) Error(msg string) {
	l.logger.Error().Msg(msg)
}

func (l *ZeroLogger) Fatal(msg string) {
	l.logger.Fatal().Msg(msg)
}
