package controller

import (
	"net/http"
	"quiz/config"
	"quiz/structs"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	var books []structs.Books
	if err := config.DB.Select("id", "title").Find(&books).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, books)
}

func AddBooks(c *gin.Context) {
	var books structs.Books

	if err := c.ShouldBindJSON(&books); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if books.TotalPage <= 100 {
		books.Thickness = "tipis"
	} else {
		books.Thickness = "tebal"
	}

	if books.ReleaseYear < 1980 || books.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book is not valid"})
	}
	if err := config.DB.Create(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Message": "New Book Added", "book": books})

}

func DetailBooks(c *gin.Context) {
	id := c.Param("id")
	var books structs.Books

	if err := config.DB.Where("id = ?", id).Find(&books).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, books)
}

func DeleteBooks(c *gin.Context) {
	id := c.Param("id")
	var books structs.Books

	if err := config.DB.First(&books, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book Not Found"})
		return
	}

	if err := config.DB.Delete(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book Deleted Succesfully",
	})

}

func GetAllCategories(c *gin.Context) {
	var category []structs.Category
	if err := config.DB.Select("id", "name").Find(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, category)
}

func AddCategories(c *gin.Context) {
	var category structs.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Message": "New Category Added", "category": category})
}

func DetailCategories(c *gin.Context) {
	id := c.Param("id")
	var category structs.Category

	if err := config.DB.Where("id = ?", id).Find(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, category)
}

func DeleteCategories(c *gin.Context) {
	id := c.Param("id")
	var category structs.Category

	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category Not Found"})
		return
	}

	if err := config.DB.Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category Deleted Succesfully",
	})
}

func GetBooksByCategory(c *gin.Context) {
	catId := c.Param("id")
	var books []structs.Books
	var category structs.Category

	if err := config.DB.First(&category, catId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category Not Found"})
		return
	}

	if err := config.DB.Where("category_id = ?", catId).Select("title").Find(&books).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book Not Found"})
		return
	}

	c.JSON(http.StatusCreated, books)
}

func AddUser(c *gin.Context) {
	var user structs.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var check structs.User

	if n := config.DB.Where("username = ?", user.Username).First(&check).Error; n == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "username existed"})
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"user":    user,
	})

}
