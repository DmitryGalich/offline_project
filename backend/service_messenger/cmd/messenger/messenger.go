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
	logger.Info("KEK2")
	logger.Info("KEK3")
	logger.Info("KEK4")
}
