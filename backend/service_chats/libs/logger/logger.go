package logger

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogger() (file *os.File) {
	folderPath := "../log/"
	filePath := folderPath + "log.txt"

	log.SetFormatter(&log.TextFormatter{})

	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
		return nil
	}

	if _, err := os.Stat(filePath); err == nil {
		log.Info("Removing old log file...")

		err = os.Remove(filePath)
		if err != nil {
			log.Fatalf("Failed to remove old log file: %v", err)
			return nil
		}

		log.Info("Old log file removed")
	}

	file, err = os.Create(filePath)
	if err == nil {
		// Create a MultiWriter to write logs to both console and file
		multiWriter := io.MultiWriter(os.Stdout, file)
		log.SetOutput(multiWriter)
	} else {
		log.Fatalf("Failed to create log file: %v", err)
	}

	return file
}
