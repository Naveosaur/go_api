package main

import (
	"fmt"
	"go_gin/config"
	"go_gin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.LoadDB()

	r := gin.Default()
	api := r.Group("/api")



	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.AuthRouter(api)
	router.PostRouter(api)

	r.Run(fmt.Sprintf(":%v", config.ENV.PORT))
}
