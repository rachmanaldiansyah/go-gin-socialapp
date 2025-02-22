package exceptions

import (
	"go-gin-sosmed/dto"
	"go-gin-sosmed/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
  var statusCode int
  
  switch err.(type) {
  case *NotFoundError:
    statusCode = http.StatusNotFound
  case *BadRequestError:
    statusCode = http.StatusBadRequest
  case *InternalServerError:
    statusCode = http.StatusInternalServerError
  case *UnauthorizedError:
    statusCode = http.StatusUnauthorized
  }
  
  response := helpers.Response(dto.ResponseParams{
    StatusCode: statusCode,
    Message: err.Error(),
  })

  c.JSON(statusCode, response)
}
