package routers

import (
	"sesi9/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/user", controller.GetUsers)
	router.GET("/user/:id", controller.GetUserById)
	router.POST("/user", controller.AddUser)

	return router
}
