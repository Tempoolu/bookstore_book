package controllers

import (
	"net/http"

	"github.com/Tempoolu/bookstore_book/models"
	"github.com/gin-gonic/gin"
)

func listBook(c *gin.Context) {
	var book models.Book
	books := book.List()
	c.JSON(http.StatusOK, gin.H{
		"result": "success",
		"book":   books,
	})
}

func getBook(c *gin.Context) {
	var book models.Book
	book.Get(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"result": "success",
		"book":   book,
	})
}

func createBook(c *gin.Context) {
	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  "failure",
			"message": err.Error(),
		})
		return
	}
	book.Create()
	c.JSON(http.StatusOK, gin.H{
		"result": "success",
	})
}

func updateBook(c *gin.Context) {
	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  "failure",
			"message": err.Error(),
		})
		return
	}
	book.Update(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"result": "success",
	})
}
