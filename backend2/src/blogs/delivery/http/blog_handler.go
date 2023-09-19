package http

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/y-maruyama1002/Techport/domain"
)

type BlogHandler struct {
	BlgUsecase domain.BlogUsecase
}

func NewBlogHandler(engine *gin.Engine, blgUcase domain.BlogUsecase) {
	handler := &BlogHandler{
		BlgUsecase: blgUcase,
	}
	engine.GET("api/v1/blogs/:id", handler.GetById)
}

func (h *BlogHandler) GetById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	blog, err := h.BlgUsecase.GetById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("cant get blog from id: %d, error is %v", id, err),
		})
	}
	c.JSON(200, blog)
}
