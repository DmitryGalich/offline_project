package websockets_manager

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var WebsocketsManager = newConnectionManager()

type ConnectionManager struct {
	connections map[string]*websocket.Conn
}

func newConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[string]*websocket.Conn),
	}
}

func (cm *ConnectionManager) AddConnection(id string, conn *websocket.Conn) {
	log.Debug("Client connected:", id)

	cm.connections[id] = conn

	log.Debug("Client connected:", id)
}

func (cm *ConnectionManager) RemoveConnection(id string) {
	log.Debug("Removing client ", id)

	delete(cm.connections, id)

	log.Debug("Client " + id + " removed")
}

func (cm *ConnectionManager) GetConnection(id string) *websocket.Conn {
	conn, ok := cm.connections[id]
	if !ok {
		return nil
	}
	return conn
}
