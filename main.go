package main

import "github.com/gin-gonic/gin"
import "github.com/itsmaheshkariya/gin-gorm-rest/routes"

func main() {
	router:= gin.New()
	routes.UserRoute(router)
	router.Run(":9004")
}
