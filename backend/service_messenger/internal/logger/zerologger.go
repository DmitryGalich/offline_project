package logger

import (
	"errors"
	"os"
	"path/filepath"
	"runtime/debug"
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

	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return nil, errors.New("can't read build info")
	}

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Int("pid", os.Getpid()).
		Str("go_version", buildInfo.GoVersion).
		Caller().
		Logger()

	fileHandler, err := lumberjack.New(
		lumberjack.WithFileName(filepath.Join(logFolderPath, logFileName)),
		lumberjack.WithMaxBytes(25*lumberjack.MB),
		lumberjack.WithMaxBackups(3),
		lumberjack.WithMaxDays(28),
		lumberjack.WithCompress(),
	)
	if err != nil {
		return nil, err
	}

	logger = logger.Output(zerolog.ConsoleWriter{Out: fileHandler})

	return &ZeroLogger{logger}, nil
}

func (l *ZeroLogger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *ZeroLogger) Warning(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *ZeroLogger) Error(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *ZeroLogger) Critical(msg string) {
	l.logger.Info().Msg(msg)
}
