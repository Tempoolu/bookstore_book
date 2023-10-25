package main

import (
	"github.com/Tempoolu/bookstore_book/controllers"
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
	models.InitDB(database.Key("user").MustString(""), database.Key("password").MustString(""), database.Key("address").MustString(""), database.Key("database").MustString(""))
	// models.Migrate()

	e := gin.Default()
	controllers.InitRouter(e)
	e.Run(":8000")
}
