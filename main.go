package main

import (
	"os"
	"quiz/config"
	"quiz/controller"
	"quiz/middleware"
	"quiz/structs"

	"github.com/gin-gonic/gin"
)

var err error

func main() {

	config.ConnectDB() //connect to database

	config.DB.AutoMigrate(&structs.Books{}, &structs.Category{}, &structs.User{}) // migrate the database

	router := gin.Default() //setup router

	router.POST("/api/signup", controller.AddUser)

	auth := router.Group("/api", middleware.BasicAuth()) //masking the api with basicauth
	{
		auth.GET("/books", controller.GetAllBooks)
		auth.POST("/books", controller.AddBooks)
		auth.GET("/books/:id", controller.DetailBooks)
		auth.DELETE("/books/:id", controller.DeleteBooks)
		auth.GET("/categories", controller.GetAllCategories)
		auth.POST("/categories", controller.AddCategories)
		auth.GET("/categories/:id", controller.DetailCategories)
		auth.DELETE("/categories/:id", controller.DeleteCategories)
		auth.GET("/categories/:id/books", controller.GetBooksByCategory)
	}
	// err = godotenv.Load("config/connect.env")
	// if err != nil {
	// 	panic("Error loading .env file")
	// }
	// port := os.Getenv("DB_PORT")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	router.Run("+", PORT)
}
