package main

import (
	"offline_project/backend/service_messenger/internal/logger"

	"path/filepath"
)

func getLogFolderPath() string {
	return filepath.Join("..", "..", "log")
}

func main() {
	logFolderPath := getLogFolderPath()
	logger, err := logger.NewZeroLogger(logFolderPath, "logfile.log")
	if err != nil {
		println(err)
		return
	}
	logger.Info("KEK1")
}
