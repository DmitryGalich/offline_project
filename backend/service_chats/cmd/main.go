// package main

// import (
// 	"io"
// 	"net/http"
// 	"os"

// 	"github.com/gorilla/websocket"
// 	log "github.com/sirupsen/logrus"
// )

// func initLogger() (file *os.File) {
// 	folderPath := "../log/"
// 	filePath := folderPath + "log.txt"

// 	log.SetFormatter(&log.TextFormatter{})

// 	err := os.MkdirAll(folderPath, 0755)
// 	if err != nil {
// 		log.Fatalf("Failed to create log directory: %v", err)
// 		return nil
// 	}

// 	if _, err := os.Stat(filePath); err == nil {
// 		log.Info("Removing old log file...")

// 		err = os.Remove(filePath)
// 		if err != nil {
// 			log.Fatalf("Failed to remove old log file: %v", err)
// 			return nil
// 		}

// 		log.Info("Old log file removed")
// 	}

// 	file, err = os.Create(filePath)
// 	if err == nil {
// 		// Create a MultiWriter to write logs to both console and file
// 		multiWriter := io.MultiWriter(os.Stdout, file)
// 		log.SetOutput(multiWriter)
// 	} else {
// 		log.Fatalf("Failed to create log file: %v", err)
// 	}

// 	return file
// }

// func getEntryPointAddress() string {
// 	entryPointAddress := os.ExpandEnv("${ENTRYPOINT_ADDRESS}:${ENTRYPOINT_PORT}")
// 	if entryPointAddress == "" {
// 		entryPointAddress = ":80"
// 	}

// 	return entryPointAddress
// }

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }
// var clients []websocket.Conn

// func processWs(w http.ResponseWriter, r *http.Request) {
// 	log.Info("Getting request")

// 	connection, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Error("Error while upgrading to websockets: ", err)
// 		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
// 		return
// 	}
// 	clients = append(clients, *connection)
// 	defer connection.Close()

// 	log.Info("Upgraded")

// 	for {
// 		messageType, message, err := connection.ReadMessage()
// 		if err != nil {
// 			log.Error("Error reading message:", err)
// 			return
// 		}

// 		log.Info("Received message: ", message)

// 		err = connection.WriteMessage(messageType, message)
// 		if err != nil {
// 			log.Error("Error echoing message:", err)
// 			return
// 		}
// 	}
// }

// func main() {
// 	logFile := initLogger()
// 	defer logFile.Close()

// 	entryPointAddress := getEntryPointAddress()
// 	log.Info("Starting server at (" + entryPointAddress + ")")

// 	http.HandleFunc("/", processWs)
// }

package main

import (
	"fmt"
	"net/http"
	"os"
	"service_chats/libs/logger"
	"service_chats/libs/websockets_manager"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	log "github.com/sirupsen/logrus"
)

func getEntryPointAddress() string {
	entryPointAddress := os.ExpandEnv("${ENTRYPOINT_ADDRESS}:${ENTRYPOINT_PORT}")
	if entryPointAddress == "" {
		log.Info("Default entry point")
		entryPointAddress = ":80"
	}

	return entryPointAddress
}

func processWs(w http.ResponseWriter, r *http.Request) {
	conn, err := websockets_manager.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}

	id := uuid.New().String()
	websockets_manager.WebsocketsManager.AddConnection(id, conn)

	defer func() {
		websockets_manager.WebsocketsManager.RemoveConnection(id)
		conn.Close()
	}()

	log.Info("Client connected:", id)

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Error("Error reading message:", err)
			}
			break
		}

		log.Info("Received message(type: " + fmt.Sprint(messageType) + ") from client " + id + " : " + string(message))

		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Error("Error echoing message:", err)
			break
		}
	}
}

func main() {
	logFile := logger.InitLogger()
	defer logFile.Close()

	log.Info("Loading entry point...")
	entryPointAddress := getEntryPointAddress()
	log.Info("Entry point - " + entryPointAddress)

	http.HandleFunc("/", processWs)

	log.Info("Server start")
	log.Fatal(http.ListenAndServe(entryPointAddress, nil))
}
