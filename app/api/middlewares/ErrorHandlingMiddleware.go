package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"goproj2/api/controllers"
)

func HandleError() gin.HandlerFunc {
	return func (c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}

		c.Errors.Last()

		// Use reflect.TypeOf(err.Err) to know the type of your error
		if error, ok := errors.Cause(err.Err).(*controllers.ParameterError); ok {
			c.JSON(c.Writer.Status(), gin.H{"error":"invalid parameter for: " + error.Error(), "details":err.Error()})
			return
		}
		if error, ok := errors.Cause(err.Err).(*controllers.ObjectNotFound); ok {
			res, _ := json.Marshal(error.Pk)
			c.JSON(c.Writer.Status(), gin.H{"error":"object with type '" + error.ObjectType.String() + "' and with pk: " + string(res) + " cannot be found", "details":error.Error()})
			return
		}
		if error, ok := errors.Cause(err.Err).(*controllers.BindError); ok {
			c.JSON(c.Writer.Status(), gin.H{"error":"error while binding model", "details":error.Error()})
			return
		}
	}
}
