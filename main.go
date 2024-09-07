package main

import (
	"ducky-dl/handler"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	router.Static("/uploads", "/app/uploads")

	validate := validator.New(validator.WithRequiredStructEnabled())

	router.POST("/api/download", func(ctx *gin.Context) {
		handler.Download(ctx, validate)
	})

	err := router.Run(":10000")
	if err != nil {
		log.Fatalf("Server failed to run!")
	}
}
