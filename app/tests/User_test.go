package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"goproj2/api/controllers"
	"goproj2/api/middlewares"
	"goproj2/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

type MockUserRepository struct {
}

func (r *MockUserRepository) Update(toUpdate *models.User, update *models.User) error {
	return nil
}

func (r *MockUserRepository) Create(user *models.User) error {
	return nil
}

func (r *MockUserRepository) Delete(user *models.User) error {
	return nil
}

func (r *MockUserRepository) GetAll() ([]models.User, error) {
	return nil, nil
}

func (r *MockUserRepository) GetById(id uint) (models.User, error) {
	if id == 10 {
		return models.User{}, errors.New("")
	}

	user := models.User{
		Name:     "test_name",
		Email:    "test@test.com",
		Password: "test_surname",
	}
	return user, nil
}

func SetupTestRouter() *gin.Engine {
	//init router
	r := gin.New()

	r.Use(middlewares.HandleError())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			if err == "UniqueConstraintFailed" {
				c.String(http.StatusForbidden, fmt.Sprintf("Already exists: %s", err))
				return
			}
			if err == "ObjectNotFound" {
				c.String(http.StatusNotFound, fmt.Sprintf("ObjectNotFound: %s", err))
				return
			}
			if err == "ArgumentException" {
				c.String(http.StatusBadRequest, fmt.Sprintf("ArgumentException: %s", err))
				return
			}
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}

		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world", "number": 3})
	})

	// Routes
	r.GET("/users/:id", func(context *gin.Context) {
		controllers.GetUser(context, &MockUserRepository{})
	})
	r.GET("/users", func(context *gin.Context) {
		controllers.ListUsers(context, &MockUserRepository{})
	})
	r.PATCH("/users/:id", func(context *gin.Context) {
		controllers.UpdateUser(context, &MockUserRepository{})
	})
	r.POST("/users", func(context *gin.Context) {
		controllers.CreateUser(context, &MockUserRepository{})
	})
	r.DELETE("/users/:id", func(context *gin.Context) {
		controllers.DeleteUser(context, &MockUserRepository{})
	})
	return r
}

func TestGetUser_Should_Return_404_When_User_Not_Found(t *testing.T) {
	router := SetupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/10", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestGetUser_Should_Return_200(t *testing.T) {
	router := SetupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/1", nil)
	router.ServeHTTP(w, req)
	var res models.User
	json.Unmarshal([]byte(w.Body.String()), &res)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "test_name", res.Name)
}

func TestUpdateUser_Should_Return_404_When_User_Not_Found(t *testing.T) {
	router := SetupTestRouter()
	data := url.Values{}
	data.Set("name", "new_name")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/users/10", strings.NewReader(data.Encode()))
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestCreateUser_Should_Return_400_When_Parameters_Are_Missing(t *testing.T) {
	router := SetupTestRouter()
	data := url.Values{}
	data.Set("name", "new_name")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(data.Encode()))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
