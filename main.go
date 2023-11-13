package main

import "github.com/gin-gonic/gin"
import "github.com/itsmaheshkariya/gin-gorm-rest/routes"
import "github.com/itsmaheshkariya/gin-gorm-rest/config"

func main() {
	router:= gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":9004")
}
