package main

import (
	"ducky-dl/handler"
	"ducky-dl/server"
	"log"
	"os"
)

func main() {
	s := server.New()

	s.Router.ForwardedByClientIP = true
	s.Router.SetTrustedProxies([]string{os.Getenv("NGINX_PROXY")})

	s.Router.Static("/uploads", "/app/uploads")

	s.Router.POST("/api/download", handler.Download)

	err := s.Router.Run(":10000")
	if err != nil {
		log.Fatalf("Server failed to run!")
	}
}
