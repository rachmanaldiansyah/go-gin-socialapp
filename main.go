package main

import (
	"fmt"
	"go-gin-sosmed/config"
	"go-gin-sosmed/routers"

	"github.com/gin-gonic/gin"
)

func main() {
  config.LoadConfig()
  config.ConnectDB()

  r := gin.Default()
  api := r.Group("/api")

  api.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })

  routers.AuthRouter(api)
  routers.PostRouter(api)

  r.Run(fmt.Sprintf(":%v", config.ENV.PORT))
}

