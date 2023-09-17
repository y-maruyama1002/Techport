package infrastructure

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func init() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1.GET("/", func(c *gin.Context) {
	  c.JSON(200, gin.H {
		"message": "pongeeeeer!!!",
	  })
	})

	Router = r
}
