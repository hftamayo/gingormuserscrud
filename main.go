package main

import "github.com/gin-gonic/gin"
import "github.com/hftamayo/gingormrest/routes"
import "github.com/hftamayo/gingormrest/config"

func main() {
	router:= gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":9004")
}
