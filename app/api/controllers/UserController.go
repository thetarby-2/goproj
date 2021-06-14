package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	. "goproj2/core"
	"goproj2/db"
	"goproj2/models"
	"net/http"
	"reflect"
	"strconv"
)

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" validate:"email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// GetObjectByRepo util method for simple crud views
func GetObjectByRepo(repository db.IUserRepository, c *gin.Context, key string, obj interface{}) error {
	id, errConvert := strconv.Atoi(c.Param(key))
	if errConvert != nil {
		c.AbortWithError(http.StatusBadRequest, &ParameterError{ParameterName: "id", Err: errConvert})
		return errConvert
	}

	res, err := repository.GetById(uint(id))
	if err != nil {
		c.AbortWithError(http.StatusNotFound, &ObjectNotFound{ObjectType: reflect.TypeOf(obj), Err: err, Pk: id})
		return err
	}
	va := reflect.ValueOf(res)
	reflect.ValueOf(obj).Elem().Set(va)

	return nil
}

// GetUser @Router /users/:id [get]
func GetUser(c *gin.Context, repository db.IUserRepository) {
	// Get model if exist
	var user models.User
	if err := GetObjectByRepo(repository, c, "id", &user); err != nil {
		return
	}
	c.JSON(http.StatusOK, user.ToUserView())
}

// CreateUser @Router /users [post]
func CreateUser(c *gin.Context, repository db.IUserRepository) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithError(http.StatusBadRequest, &BindError{Err: err})
		return
	}
	validate := validator.New()

	if err := validate.Struct(input); err != nil {
		c.AbortWithError(http.StatusBadRequest, &ValidationError{Err: err})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(passwordHash),
	}

	if err := repository.Create(&user); err != nil {
		c.AbortWithError(403, &DbConstraintCheckFailed{Name: "UniqueConstraint", Err: err})
		return
	}

	c.JSON(http.StatusOK, user.ToUserView())
}

// ListUsers @Router /users [get]
func ListUsers(c *gin.Context, repository db.IUserRepository) {
	// Get model if exist
	users, err := repository.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var userViews = make([]models.UserView, 0)
	for _, user := range users {
		userViews = append(userViews, user.ToUserView())
	}

	c.JSON(http.StatusOK, userViews)
}

// UpdateUser @Router /users/:id [patch]
func UpdateUser(c *gin.Context, repository db.IUserRepository) {
	// Get model if exist
	var user models.User
	if err := GetObjectByRepo(repository, c, "id", &user); err != nil {
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

	c.JSON(http.StatusOK, user.ToUserView())
}

// DeleteUser @Router /users/:id [delete]
func DeleteUser(c *gin.Context, repository db.IUserRepository) {
	// Get model if exist
	var user models.User
	if err := GetObjectByRepo(repository, c, "id", &user); err != nil {
		return
	}

	repository.Delete(&user)

	c.JSON(http.StatusNoContent, "")
}
