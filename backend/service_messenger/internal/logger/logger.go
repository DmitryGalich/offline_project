package logger

import (
	"os"
)

type Logger interface {
	Info(msg string)
	Warning(msg string)
	Error(msg string)
	Critical(msg string)
}

func prepareLogFolder(logFolderPath string) error {
	err := os.RemoveAll(logFolderPath)
	if err != nil {
		return err
	}

	err = os.MkdirAll(logFolderPath, 0755)
	if err != nil {
		return err
	}

	return nil
}
