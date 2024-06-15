package handlers

import (
	"crud_app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.HTML(http.StatusOK, "index.html", gin.H{"books": books})
}

func ShowNewBookForm(c *gin.Context) {
	c.HTML(http.StatusOK, "new.html", nil)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBind(&book); err == nil {
		models.DB.Create(&book)
		c.Redirect(http.StatusFound, "/")
	} else {
		c.HTML(http.StatusBadRequest, "layout.html", gin.H{"error": err.Error()})
	}
}

func ShowEditBookForm(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book
	models.DB.First(&book, id)
	c.HTML(http.StatusOK, "edit.html", gin.H{"book": book})
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book
	if err := c.ShouldBind(&book); err == nil {
		models.DB.Model(&book).Where("id = ?", id).Updates(book)
		c.Redirect(http.StatusFound, "/")
	} else {
		c.HTML(http.StatusBadRequest, "edit.html", gin.H{"error": err.Error(), "book": book})
	}
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	models.DB.Delete(&models.Book{}, id)
	c.Redirect(http.StatusFound, "/")
}
