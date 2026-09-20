package routes

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	route := gin.Default()
	InitRoute(route)
	route.Run(":8080")
}
