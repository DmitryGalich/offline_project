package websockets_manager

import (
	"github.com/gorilla/websocket"
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
	cm.connections[id] = conn
}

func (cm *ConnectionManager) RemoveConnection(id string) {
	delete(cm.connections, id)
}

func (cm *ConnectionManager) GetConnection(id string) *websocket.Conn {
	conn, ok := cm.connections[id]
	if !ok {
		return nil
	}
	return conn
}
