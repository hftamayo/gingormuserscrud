package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/itsmaheshkariya/gin-gorm-rest/controller"
)

func UserRoute(router *gin.Engine){
	router.GET("/view-users", controller.GetUsers)
	router.POST("/add-user", controller.CreateUser)
	router.DELETE("/delete-user/:id", controller.DeleteUser)
	router.PUT("/update-user/:id", controller.UpdateUser)
}
