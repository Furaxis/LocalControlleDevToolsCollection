package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListaPostagens(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "OLA, OLA, OLA, OLA",
	})
}
