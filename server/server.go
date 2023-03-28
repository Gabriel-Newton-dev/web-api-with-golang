package server

import (
	"log"

	"github.com/Gabriel-Newton-dev/web-api-with-golang/router"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := router.ConfigRoutes(s.server)

	log.Print("Server is running at port:", s.port)
	log.Fatal(router.Run(":" + s.port))
}
