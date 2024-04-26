package server

import (
	"fmt"
	"net/http"
	"offline_project/backend/service_messenger/internal/logger"
)

type BasicServer struct {
	logger logger.Logger
	config ServerConfig
}

func NewBasicServer(logger logger.Logger, config ServerConfig) *BasicServer {
	return &BasicServer{logger: logger, config: config}
}

func (s *BasicServer) handleHome(w http.ResponseWriter, req *http.Request) {
	s.logger.Info("\n" +
		"req.Host: " + req.Host + "\n" +
		"req.Method: " + req.Method + "\n" +
		"req.RemoteAddr: " + req.RemoteAddr + "\n" +
		"req.RequestURI: " + req.RequestURI + "\n")

	fmt.Fprintf(w, "home")
}

func (s *BasicServer) Start() error {
	s.logger.Info("start basicServer Address: " + s.config.address +
		" port: " + s.config.port,
	)

	http.HandleFunc("/", s.handleHome)
	return http.ListenAndServe(s.config.address+":"+s.config.port, nil)
}
