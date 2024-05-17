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

func (s *BasicServer) handleTest(w http.ResponseWriter, req *http.Request) {
	anwser := "test"

	s.logger.Info("\n" +
		"req.Host: " + req.Host + "\n" +
		"req.Method: " + req.Method + "\n" +
		"req.RemoteAddr: " + req.RemoteAddr + "\n" +
		"req.RequestURI: " + req.RequestURI + "\n" +
		"answer: " + anwser + "\n")

	fmt.Fprint(w, anwser)
}

func (s *BasicServer) handleWs(w http.ResponseWriter, req *http.Request) {
	anwser := "ws"

	s.logger.Info("\n" +
		"req.Host: " + req.Host + "\n" +
		"req.Method: " + req.Method + "\n" +
		"req.RemoteAddr: " + req.RemoteAddr + "\n" +
		"req.RequestURI: " + req.RequestURI + "\n" +
		"answer: " + anwser + "\n")

	fmt.Fprint(w, anwser)
}

func (s *BasicServer) Start() error {
	s.logger.Info("start basicServer Address: " + s.config.address +
		" port: " + s.config.port,
	)

	http.HandleFunc("/test/", s.handleTest)
	http.HandleFunc("/ws/", s.handleWs)

	return http.ListenAndServe(s.config.address+":"+s.config.port, nil)
}
