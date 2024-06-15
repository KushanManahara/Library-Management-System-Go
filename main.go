package main

import (
	"crud_app/handlers"
	"crud_app/models"
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	models.InitDatabase()

	// Initialize the router
	router := gin.Default()
	router.Static("/static", "./static")
	router.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	router.LoadHTMLGlob("templates/*")

	// Define the routes
	router.GET("/", handlers.ShowBooks)
	router.GET("/books/new", handlers.ShowNewBookForm)
	router.GET("/books/delete/:id", handlers.DeleteBook)
	router.GET("/books/edit/:id", handlers.ShowEditBookForm)

	router.POST("/books", handlers.CreateBook)
	router.POST("/books/update/:id", handlers.UpdateBook)

	// Start the server
	router.Run(":8080")
}
