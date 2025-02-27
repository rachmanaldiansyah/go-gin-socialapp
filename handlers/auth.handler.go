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

func (h *authHandler) Login(c *gin.Context) {
  var login dto.LoginRequest
  
  err := c.ShouldBindJSON(&login)
  if err != nil {
    exceptions.HandleError(c, &exceptions.BadRequestError{Message: err.Error()})
    return
  }

  result, err := h.services.Login(&login)
  if err != nil {
    exceptions.HandleError(c, err)
    return
  }

  res := helpers.Response(dto.ResponseParams{
    StatusCode: http.StatusOK,
    Message:    "Login Successfully",
    Data:       result,
  })
  
  c.JSON(http.StatusOK, res)
}
