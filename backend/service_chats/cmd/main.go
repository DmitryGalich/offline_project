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
	log.Info("Upgrading connection...")

	conn, err := websockets_manager.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}

	id := uuid.New().String()
	websockets_manager.WebsocketsManager.AddConnection(id, conn)

	defer func() {
		log.Info("Closing connection " + id)
		websockets_manager.WebsocketsManager.RemoveConnection(id)
		conn.Close()
		log.Info("Connection " + id + "  closed")
	}()

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
