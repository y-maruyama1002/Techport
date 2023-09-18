package main

import (
	"github.com/gin-gonic/gin"
	"github.com/y-maruyama1002/Techport/router"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title  string
	Body string
  }

func main() {
	engine := gin.Default()

	dsn := "root:password@tcp(db)/root?charset=utf8mb4&parseTime=True&loc=Local"
  	dbEngine, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	dbEngine.AutoMigrate(&Blog{})

	var blog Blog
	dbEngine.First(&blog, 3)

	router.SetRoutes(engine, dbEngine)
	engine.Run(":3000")
}
