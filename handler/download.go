package handler

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"regexp"

	"github.com/gin-gonic/gin"
)

type ClientRequest struct {
	Link string `json:"link"`
}

type ServerResponse struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	URL      string `json:"url"`
	Filename string `json:"filename"`
}

func Download(ctx *gin.Context) {
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

	// TODO: Refactor and optimize in the way that you don't have to recompile the pattern at each request
	pattern := `(?:https?:\/\/)?(?:www\.)?youtu(?:\.be\/|be\.com\/(?:watch\?v=|embed\/|v\/|.*[?&]v=))?([-a-zA-Z0-9_]{11})`
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Printf("Error: Invalid link - %v, %v", req.Link, err)
		ctx.JSON(
			http.StatusInternalServerError,
			ServerResponse{
				Success:  false,
				Message:  "Internal server error",
				URL:      "",
				Filename: "",
			},
		)
		return
	}

	matches := re.MatchString(req.Link)
	if !matches {
		log.Printf("Error: Invalid link - %v, %v", req.Link, err)
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
}
