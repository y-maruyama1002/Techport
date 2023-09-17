package main

import (
	"fmt"

	"github.com/y-maruyama1002/Techport/infrastructure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
  }

func main() {
	dsn := "root:password@tcp(db)/root?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	  // Migrate the schema
	  db.AutoMigrate(&Product{})

	  // Create
	  db.Create(&Product{Code: "D43", Price: 200})

	  var product Product
	  db.First(&product, 1)
	  fmt.Println("check the value")
	  fmt.Println(product.Code)
	  //  D42
	  fmt.Println(product.Price)
	  // 100

	infrastructure.Router.Run(":3000")
}
