package router

import (
	"go_gin/config"
	"go_gin/handler"
	"go_gin/middleware"
	"go_gin/repository"
	"go_gin/service"

	"github.com/gin-gonic/gin"
)

func PostRouter(api *gin.RouterGroup){
	postRepository := repository.NewPostRepository(config.DB)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	r := api.Group("/posts")

	r.Use(middleware.JWTMiddleware())

	

	r.POST("/", postHandler.Create)
}