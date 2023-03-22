package server

import (
	"fmt"
	"fx-bank/middlewares"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
)

type Server struct {
	e   *gin.Engine
	srv http.Server
}

func New() *Server {
	return &Server{
		e: gin.Default(),
	}
}

func (s *Server) SetupMiddlewares(m []gin.HandlerFunc) {
	s.e.Use(m...)
}

// TODO: Set up route groups when we have authenticated routes and pass JWT middleware to authenticated routes
func (s *Server) SetupRoutes() *gin.Engine {
	mw := []gin.HandlerFunc{middlewares.Cors()}
	s.SetupMiddlewares(mw)
	return s.e
}

func (s *Server) Start() {

	s.srv = http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: s.SetupRoutes(),
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		if err := s.srv.Close(); err != nil {
			log.Println("failed to shutdown server", err)
		}
	}()

	if err := s.srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("server closed after interruption")
		} else {
			log.Println("unexpected server shutdown. err:", err)
		}
	}
}

func (s *Server) Stop() error {
	return s.srv.Close()
}
