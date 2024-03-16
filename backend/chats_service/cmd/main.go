package main

import (
	"os"

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
	file, err = os.Create(filePath)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Fatalf("Failed to create log file: %v", err)
	}

	return file
}

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// func getEntryPointAddress() string {
// 	entryPointAddress := os.ExpandEnv("${ENTRYPOINT_ADDRESS}:${ENTRYPOINT_PORT}")
// 	if entryPointAddress == "" {
// 		entryPointAddress = ":80"
// 	}

// 	return entryPointAddress
// }

// func home(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintf(w, "Chats service - Home page")
// }

// func websockets(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
// 		return
// 	}
// 	defer conn.Close()

// 	for {
// 		messageType, message, err := conn.ReadMessage()
// 		if err != nil {
// 			// fmt.Println("Error reading message:", err)
// 			return
// 		}

// 		// fmt.Printf("Received message: %s\n", message)

// 		// Echo back the received message
// 		err = conn.WriteMessage(messageType, message)
// 		if err != nil {
// 			// fmt.Println("Error echoing message:", err)
// 			return
// 		}
// 	}
// }

func main() {
	logFile := initLogger()
	defer logFile.Close()

	log.Info("LEL")

	// entryPointAddress := getEntryPointAddress()
	// log..Println("Starting server at: \"" + entryPointAddress + "\"")

	// http.HandleFunc("/", home)
	// http.HandleFunc("/ws", websockets)
	// log "github.com/sirupsen/logrus"

}
