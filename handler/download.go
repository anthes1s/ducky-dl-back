package handler

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ClientRequest struct {
	Link string `json:"link" validate:"url"`
}

type ServerResponse struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	URL      string `json:"url"`
	Filename string `json:"filename"`
}

func Download(ctx *gin.Context, validate *validator.Validate) {
	var req ClientRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		log.Printf("Error occurred - %v", err)
		ctx.JSON(
			http.StatusBadRequest,
			ServerResponse{
				Success:  false,
				Message:  "Error occured while downloading a file",
				URL:      "",
				Filename: "",
			},
		)
		return
	}

	err = validate.Struct(&req)
	if err != nil {
		log.Printf("Error occurred - %v", err)
		ctx.JSON(
			http.StatusBadRequest,
			ServerResponse{
				Success:  false,
				Message:  "Invalid link",
				URL:      "",
				Filename: "",
			},
		)
		return
	}
	cmd := exec.Command(
		"yt-dlp",
		"-f", "mp4",
		"--get-filename",
		req.Link,
	)
	filename, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error occurred - %v", err)
		ctx.JSON(
			http.StatusBadRequest,
			ServerResponse{
				Success:  false,
				Message:  "Error occured while downloading a file, failed to get a files name",
				URL:      "",
				Filename: "",
			},
		)
		return
	}

	// TODO: Find a way to constantly sent progress% to the user
	// 1. Use WebSockets
	// 2. Send standard output to the user
	// 3. Render progress in the client
	cmd = exec.Command(
		"yt-dlp",
		"-f", "mp4",
		"-P", "/app/uploads",
		"--progress",
		"--no-playlist",
		req.Link,
	)
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error occurred - %v", err)
		ctx.JSON(
			http.StatusBadRequest,
			ServerResponse{
				Success:  false,
				Message:  "Error occured while downloading a file, failed to download a file",
				URL:      "",
				Filename: "",
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		ServerResponse{
			Success:  true,
			Message:  "File successfully downloaded",
			URL:      fmt.Sprintf("/uploads/%s", filename),
			Filename: string(filename),
		},
	)
	return
}
