package handlers

import (
	"go-gin-sosmed/dto"
	"go-gin-sosmed/exceptions"
	"go-gin-sosmed/helpers"
	"go-gin-sosmed/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
  services services.AuthService
}

func NewAuthHandler(s services.AuthService) *authHandler {
  return &authHandler{
    services: s,
  }
}

func (h *authHandler) Register(c *gin.Context) {
  var register dto.RegisterRequest

  if err := c.ShouldBindJSON(&register); err != nil {
    exceptions.HandleError(c, &exceptions.BadRequestError{Message: err.Error()})
    return
  }

  if err := h.services.Register(&register); err != nil {
    exceptions.HandleError(c, err)
    return
  }

  res := helpers.Response(dto.ResponseParams {
    StatusCode: http.StatusCreated,
    Message: "Register Successfully, please login",
  })

  c.JSON(http.StatusCreated, res)
}
