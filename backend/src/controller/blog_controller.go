package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/y-maruyama1002/Techport/usecase"
)

type IBlogController interface {
	GetBlog(c *gin.Context)
}

type BlogController struct {
	usecase.BlogInputPort
}

func NewBlogController(srv usecase.BlogInputPort) IBlogController {
	return &BlogController{srv}
}

func (h *BlogController) GetBlog(c *gin.Context) {
	blogId := c.Param("id")
	blog, _ := h.BlogInputPort.GetBlog(blogId)
	c.JSON(200, blog)
}
