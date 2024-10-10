// TODO: Switch to stdlib

package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func New() *Server {
	server := Server{}

	server.Router = gin.Default()

	return &server
}
