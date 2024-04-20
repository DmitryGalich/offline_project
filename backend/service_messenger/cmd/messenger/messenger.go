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
	logger.Debug("KEK1")
	logger.Info("KEK2")
	logger.Warning("KEK3")
	logger.Error("KEK4")
	logger.Fatal("KEK5")
}
