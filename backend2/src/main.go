package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Blog struct {
  ID int
  Title string
  Body string
}

func main() {
	fmt.Println("hello world")

  db, err := sql.Open("mysql", "root:password@tcp(db)/root?charset=utf8mb4&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
    panic("failed to connect database")
  }
  defer db.Close()

  rows, err := db.Query("SELECT id, title, body FROM blogs")
  if err != nil {
    log.Fatalf("get rows error:%v", err)
  }
  defer rows.Close()

  for rows.Next() {
    bg := &Blog{}
    rows.Scan(&bg.ID, &bg.Title, &bg.Body)
    fmt.Println(bg)
  }

	r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pppppppp",
    })
  })
  r.Run()
}
