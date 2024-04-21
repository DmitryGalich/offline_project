package main

import (
	"offline_project/backend/service_messenger/internal/logger"
	"strconv"

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

	for j := 0; j < 1000000; j++ {
		logger.Info("KEK " + strconv.Itoa(j))
	}
}
