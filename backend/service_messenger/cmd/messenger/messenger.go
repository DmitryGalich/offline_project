package main

import (
	"offline_project/backend/service_messenger/internal/logger"
	"offline_project/backend/service_messenger/internal/server"

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

	logger.Info("Starting...")

	serverConf := server.NewServerConfig("", "80")
	server := server.NewBasicServer(logger, serverConf)

	err = server.Start()
	if err != nil {
		logger.Error("Can't start server")
	}
}
