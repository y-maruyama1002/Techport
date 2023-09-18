package main

import (
	"fmt"

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

	// Create
	// dbEngine.Create(&Blog{Title: "this is title", Body: "this is body"})

	var blog Blog
	dbEngine.First(&blog, 3)
	fmt.Println("check the value")
	fmt.Println(blog.Title)
	fmt.Println(blog.Body)

	router.SetRoutes(engine, dbEngine)
	engine.Run(":3000")
}
