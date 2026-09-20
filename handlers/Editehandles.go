package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditarPostagem(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": " Hello World ! ",
	})
}
