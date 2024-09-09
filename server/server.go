package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	Router    *gin.Engine
	Validator *validator.Validate
}

func New() *Server {
	server := Server{}

	server.Router = gin.Default()
	server.Validator = validator.New(validator.WithRequiredStructEnabled())

	return &server
}
