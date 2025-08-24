package main

import (
	"quiz/config"
	"quiz/controller"
	"quiz/middleware"
	"quiz/structs"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB() //connect to database

	config.DB.AutoMigrate(&structs.Books{}, &structs.Category{}, &structs.User{}) // migrate the database

	router := gin.Default() //setup router

	auth := router.Group("/api", middleware.BasicAuth()) //masking the api with basicauth
	{
		auth.GET("/books", controller.GetAllBooks)
		auth.POST("/books", controller.AddBooks)
		auth.GET("/books/:id", controller.DetailBooks)
		auth.DELETE("/books/:id", controller.DeleteBooks)
		auth.GET("/categories", controller.GetAllCategories)
		auth.POST("/categories", controller.AddCategories)
		auth.GET("/categories/:id", controller.DetailCategories)
		auth.DELETE("/categores/:id", controller.DeleteCategories)
		auth.GET("/categories/:id", controller.GetBooksByCategory)
	}

}
