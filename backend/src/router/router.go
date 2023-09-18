package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/y-maruyama1002/Techport/common/di"
	"gorm.io/gorm"
)

var Router *gin.Engine

func SetRoutes(engine *gin.Engine, db *gorm.DB) {
	v1 := engine.Group("/api/v1")

	blog := di.InitBlog(db)
	fmt.Println(blog)
	v1.GET("/blog/:id", blog.GetBlog)
}
