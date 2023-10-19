package main

import (
	"net/http"

	"github.com/Tempoolu/bookstore_book/models"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File
)

func main() {
	Cfg, _ = ini.Load("conf/book.ini")

	database, _ := Cfg.GetSection("database")
	models.InitDB(database.Key("USER").MustString(""), database.Key("PASSWORD").MustString(""), database.Key("ADDR").MustString(""), database.Key("DATABASE").MustString(""))
	// models.Migrate()

	r := gin.Default()
	r.GET("/book", func(c *gin.Context) {
		var book models.Book
		book.First()
		c.JSON(http.StatusOK, gin.H{
			"result": "success",
			"book":   book,
		})
	})
	r.POST("/book", func(c *gin.Context) {
		var book models.Book
		book.Create("hello_world", "some thing")
		c.JSON(http.StatusOK, gin.H{
			"result": "success",
		})
	})
	r.Run(":8000")
}
