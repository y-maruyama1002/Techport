package router

import (
	"github.com/gin-gonic/gin"
	"github.com/y-maruyama1002/Techport/common/di"
	"gorm.io/gorm"
)

var Router *gin.Engine

func SetRoutes(engine *gin.Engine, db *gorm.DB) {
	v1 := engine.Group("/api/v1")

	blog := di.InitBlog(db)
	v1.GET("/blog/:id", blog.GetBlog)
	v1.POST("/blogs", blog.CreateBlog)
}
