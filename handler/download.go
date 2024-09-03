package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Download(ctx *gin.Context) {

	ctx.JSON(http.StatusNotImplemented, "Endpoint works, but the functionality is still in work")
}
