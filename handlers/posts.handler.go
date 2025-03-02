package handlers

import (
	"go-gin-sosmed/dto"
	"go-gin-sosmed/exceptions"
	"go-gin-sosmed/helpers"
	"go-gin-sosmed/services"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type postHandler struct {
	services services.PostService
}

func NewPostHandler(services services.PostService) *postHandler {
	return &postHandler{
		services: services,
	}
}

func (h *postHandler) Create(c *gin.Context) {
	var post dto.PostRequest

	if err := c.ShouldBind(&post); err != nil {
		exceptions.HandleError(c, &exceptions.BadRequestError{Message: err.Error()})
		return
	}

	if post.PictureUrl != nil {
		if err := os.MkdirAll("/public/picture", 0755); err != nil {
			exceptions.HandleError(c, &exceptions.InternalServerError{Message: err.Error()})
			return
		}

		// rename picture
		ext := filepath.Ext(post.PictureUrl.Filename)
		newFileName := uuid.New().String() + ext
		
		// save image to directory
		dst := filepath.Join("public/picture", filepath.Base(newFileName))
		c.SaveUploadedFile(post.PictureUrl, dst)
	}

	userID := 1
	post.UserID = userID
	
	if err := h.services.Create(&post); err != nil {
		exceptions.HandleError(c, err)
		return
	}

	res := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message: "Your tweet has been posted.",
	})

	c.JSON(http.StatusCreated, res)
}