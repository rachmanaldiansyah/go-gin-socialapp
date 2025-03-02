package routers

import (
	"go-gin-sosmed/config"
	"go-gin-sosmed/handlers"
	"go-gin-sosmed/models"
	"go-gin-sosmed/services"

	"github.com/gin-gonic/gin"
)

func PostRouter(api *gin.RouterGroup) {
	postRepository := models.NewPostRepository(config.DB)
	postService := services.NewPostService(postRepository)
	postHandler := handlers.NewPostHandler(postService)

	r := api.Group("/tweets")

	r.POST("/", postHandler.Create)
}