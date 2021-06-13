package controllers

import (
	"github.com/gin-gonic/gin"
	. "goproj2/core"
	"goproj2/db"
	"goproj2/models"
	"net/http"
	"reflect"
	"strconv"
)

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// GetObjectByRepo util method for simple crud views
func GetObjectByRepo(c *gin.Context, key string, repository db.Repository, obj interface{}) error {
	id, errConvert := strconv.Atoi(c.Param(key))
	if errConvert != nil {
		c.AbortWithError(http.StatusBadRequest, &ParameterError{ParameterName: "id", Err: errConvert})
		return errConvert
	}

	err := repository.GetById(uint(id), obj)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, &ObjectNotFound{ObjectType: reflect.TypeOf(obj), Err: err, Pk: id})
		return err
	}
	return nil
}

func GetUser(c *gin.Context, repository db.Repository) {
	// Get model if exist
	var user models.User
	if err := GetObjectByRepo(c, "id", repository, &user); err != nil {
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context, repository db.Repository) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithError(http.StatusBadRequest, &BindError{Err: err})
		return
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := repository.Create(&user); err != nil {
		c.AbortWithError(403, &DbConstraintCheckFailed{Name: "UniqueConstraint", Err: err})
		return
	}

	c.JSON(http.StatusOK, user)
}

func ListUsers(c *gin.Context, repository db.Repository) {
	// Get model if exist
	users, err := repository.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, users)
}

func UpdateUser(c *gin.Context, repository db.Repository) {
	// Get model if exist
	var user models.User
	if err := GetObjectByRepo(c, "id", repository, &user); err != nil {
		return
	}

	// Validate input
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithError(http.StatusBadRequest, &BindError{Err: err})
		return
	}

	if err := repository.Update(&user, &models.User{Name: input.Name, Password: input.Password}); err != nil {
		// do sth.
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context, repository db.Repository) {
	// Get model if exist
	var user models.User
	if err := GetObjectByRepo(c, "id", repository, &user); err != nil {
		return
	}

	repository.Delete(&user)

	c.JSON(http.StatusNoContent, "")
}
