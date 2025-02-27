package routers

import (
	"go-gin-sosmed/config"
	"go-gin-sosmed/handlers"
	"go-gin-sosmed/repository"
	"go-gin-sosmed/services"

	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup) {
  authRepository := repository.NewAuthRepository(config.DB)
  authService := services.NewAuthService(authRepository)
  authHandler := handlers.NewAuthHandler(authService)

  api.POST("/register", authHandler.Register)
  api.POST("/login", authHandler.Login)
}
