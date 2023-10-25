package controllers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	engine.GET("/book", listBook)
	engine.GET("/book/:id", getBook)
	engine.POST("/book", createBook)
	engine.PUT("/book/:id", updateBook)
}
