package api

import (
	"github.com/gin-gonic/gin"
	"goproj2/api/controllers"
	"goproj2/api/middlewares"
	"goproj2/db"
)

func SetupRouter() *gin.Engine {
	//init router
	r := gin.New()

	r.Use(middlewares.HandleError())

	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Routes
	r.GET("/users/:id", func(context *gin.Context) {
		controllers.GetUser(context, &db.UserRepository{DB: db.DB})
	})
	r.GET("/users", func(context *gin.Context) {
		controllers.ListUsers(context, &db.UserRepository{DB: db.DB})
	})
	r.PATCH("/users/:id", func(context *gin.Context) {
		controllers.UpdateUser(context, &db.UserRepository{DB: db.DB})
	})
	r.POST("/users", func(context *gin.Context) {
		controllers.CreateUser(context, &db.UserRepository{DB: db.DB})
	})
	r.DELETE("/users/:id", func(context *gin.Context) {
		controllers.DeleteUser(context, &db.UserRepository{DB: db.DB})
	})
	return r
}
