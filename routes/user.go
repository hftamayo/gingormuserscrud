package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/itsmaheshkariya/gin-gorm-rest/controller"
)

func UserRoute(router *gin.Engine){
	router.GET("/", controller.UserController)
}
