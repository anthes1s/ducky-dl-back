package main

import (
	"ducky-dl/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/api/download", handler.Download)

	err := router.Run(":5000")
	if err != nil {
		log.Fatalf("Server failed to run!")
	}
}
