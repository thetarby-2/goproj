package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goproj2/api/controllers"
	"goproj2/db"
	"net/http"
)

func main() {
	// Connect to database
	db.ConnectDatabase()

	//init router
	r := gin.New()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			if err == "UniqueConstraintFailed"{
				c.String(http.StatusForbidden, fmt.Sprintf("Already exists: %s", err))
				return
			}
			if err == "ObjectNotFound"{
				c.String(http.StatusNotFound, fmt.Sprintf("ObjectNotFound: %s", err))
				return
			}
			if err == "ArgumentException"{
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
		c.JSON(http.StatusOK, gin.H{"data": "hello world", "number":3})
	})

	// Routes
	r.GET("/books/:id", controllers.GetUser)
	r.GET("/users", controllers.ListUsers)
	r.PATCH("/books/:id", controllers.UpdateUser)
	r.POST("/books", controllers.CreateUser)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run()
}
