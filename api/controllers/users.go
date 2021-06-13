package controllers

import (
	"github.com/gin-gonic/gin"
	"goproj2/db"
	"goproj2/models"
	"net/http"
	"strconv"
)

type CreateUserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	Name  string `json:"name"`
	Password string `json:"password"`
}


func GetUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	id, errConvert := strconv.Atoi(c.Param("id"))
	if errConvert != nil{
		panic("ArgumentException")
	}
	if err := db.DB.First(&user, id).Error; err != nil {
		//c.AbortWithError(404, err)
		panic("ObjectNotFound")
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}


func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := db.DB.Create(&book).Error; err != nil{
		panic("UniqueConstraintFailed")
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}


func ListUsers(c *gin.Context) {
	// Get model if exist
	var users []models.User
	db.DB.Find(&users)

	if res := db.DB.Find(&users); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}


func UpdateUser(c *gin.Context) {
	// Get model if exist
	var user models.User

	id, errConvert := strconv.Atoi(c.Param("id"))
	if errConvert != nil{
		panic("ArgumentException")
	}

	if err := db.DB.First(&user, id).Error; err != nil {
		//c.AbortWithError(404, err)
		panic("ObjectNotFound")
	}

	// Validate input
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}