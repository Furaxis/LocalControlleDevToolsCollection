package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeletarPostagem(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": " Hello World ! ",
	})
}
