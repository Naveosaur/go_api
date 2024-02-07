package handler

import (
	"fmt"
	"go_gin/dto"
	errorhandler "go_gin/errorHandler"
	"go_gin/helper"
	"go_gin/service"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type postHandler struct {
	service service.PostService
}

func NewPostHandler(service service.PostService) *postHandler{
	return &postHandler {
		service: service,
	}
}



func (h *postHandler) Create (c *gin.Context) {
	var post dto.PostRequest

	if err := c.ShouldBind(&post); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}


	if post.Picture != nil {
		if err := os.MkdirAll("/public/picture", 0755); err != nil {
			errorhandler.HandleError(c, &errorhandler.InternalServerError{Message: err.Error()})
			return
		}

		// Rename Pic
		ext := filepath.Ext(post.Picture.Filename)
		newFileName := uuid.New().String() + ext


		// Save img to directory
		dst := filepath.Join("public/picture", filepath.Base(newFileName))
		c.SaveUploadedFile(post.Picture, dst)

		post.Picture.Filename = fmt.Sprintf("%s/public/picture/%s", c.Request.Host, newFileName )
	}

	userID, _:= c.Get("userID")
	post.UserID = userID.(int) // Parsing ke int
	
	if err := h.service.Create(&post); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message: "Posted",
	})

	c.JSON(http.StatusCreated, res)

	
}