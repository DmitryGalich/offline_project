package server

import (
	"offline_project/backend/service_messenger/internal/logger"

	"github.com/panjf2000/gnet/v2"
)

type GnetServer struct {
	logger logger.Logger
	config ServerConfig
}

func NewGnetServer(logger logger.Logger, config ServerConfig) *GnetServer {
	return &GnetServer{logger: logger, config: config}
}

func (s *GnetServer) Start() error {
	s.logger.Info("start GnetServer Address: " + s.config.address +
		" port: " + s.config.port,
	)

	wss := &WsServer{addr: "tcp://" + s.config.address + ":" + s.config.port, multicore: true}
	return gnet.Run(wss, wss.addr, gnet.WithMulticore(true), gnet.WithReusePort(true), gnet.WithTicker(true))
}
