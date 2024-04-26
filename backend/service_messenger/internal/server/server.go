package server

type Server interface {
	Start()
}

type ServerConfig struct {
	address, port string
}

func NewServerConfig(address string, port string) ServerConfig {
	return ServerConfig{address: address, port: port}
}
