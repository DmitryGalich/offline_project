package main

import (
	"io"
	"net/http"
	"os"
	"fmt"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func initLogger() (file *os.File) {
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

func getEntryPointAddress() string {
	entryPointAddress := os.ExpandEnv("${ENTRYPOINT_ADDRESS}:${ENTRYPOINT_PORT}")
	if entryPointAddress == "" {
		entryPointAddress = ":80"
	}

	return entryPointAddress
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Chats service - Home page")
}

// func home(w http.ResponseWriter, r *http.Request) {
// 	log.Info("Getting request")

// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
// 		return
// 	}
// 	defer conn.Close()

// 	log.Info("Upgraded")

// 	for {
// 		messageType, message, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Error("Error reading message:", err)
// 			return
// 		}

// 		log.Info("Received message: ", message)

// 		err = conn.WriteMessage(messageType, message)
// 		if err != nil {
// 			log.Error("Error echoing message:", err)
// 			return
// 		}
// 	}
// }

func main() {
	logFile := initLogger()
	defer logFile.Close()

	entryPointAddress := getEntryPointAddress()
	log.Info("Starting server at (" + entryPointAddress + ")")

	http.HandleFunc("/", home)

	if err := http.ListenAndServe(entryPointAddress, nil); err != nil {
		log.Panic("Can't start server at (" + entryPointAddress + ")")
	}
}
