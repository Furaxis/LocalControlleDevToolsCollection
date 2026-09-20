package routes

import (
	handler "github.com/Furaxis/LocalControlleDevToolsCollection/handlers"
	"github.com/gin-gonic/gin"
)

func InitRoute(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/PostOpening", handler.Novapostagen)
		v1.GET("/PostOpening", handler.MostrarDataPostagem)
		v1.DELETE("/PostOpening", handler.DeletarPostagem)
		v1.PUT("/PostOpening", handler.EditarPostagem)
		v1.GET("/PostOpenings", handler.ListaPostagens)
	}
}
