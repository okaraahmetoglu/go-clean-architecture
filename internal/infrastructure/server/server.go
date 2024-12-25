package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	server *http.Server
	router *gin.Engine
}

func NewHTTPServer(port int, router *gin.Engine) *HTTPServer {
	return &HTTPServer{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
		router: router,
	}
}

func (s *HTTPServer) Start() error {
	log.Printf("Sunucu başlatılıyor... Port: %s", s.server.Addr)
	return s.server.ListenAndServe()
}
