package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	. "goproj2/core"
)

func HandleError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}

		c.Errors.Last()

		// Use reflect.TypeOf(err.Err) to know the type of your error
		if err_, ok := errors.Cause(err.Err).(*ParameterError); ok {
			c.JSON(c.Writer.Status(), gin.H{"error": "invalid parameter for: " + err_.Error(), "details": err.Error()})
			return
		}
		if err_, ok := errors.Cause(err.Err).(*ObjectNotFound); ok {
			res, _ := json.Marshal(err_.Pk)
			c.JSON(c.Writer.Status(), gin.H{"error": "object with type '" + err_.ObjectType.String() + "' and with pk: " + string(res) + " cannot be found", "details": err_.Error()})
			return
		}
		if err_, ok := errors.Cause(err.Err).(*BindError); ok {
			c.JSON(c.Writer.Status(), gin.H{"error": "error while binding model", "details": err_.Error()})
			return
		}
		if err_, ok := errors.Cause(err.Err).(*DbConstraintCheckFailed); ok {
			c.JSON(c.Writer.Status(), gin.H{"error": err_.Error(), "details": err_.Err.Error()})
			return
		}
		if err_, ok := errors.Cause(err.Err).(*ValidationError); ok {
			c.JSON(c.Writer.Status(), gin.H{"error": "Validation error", "details": err_.Error()})
			return
		}
		c.JSON(c.Writer.Status(), gin.H{"error": "An unhandled error occurred"})
		return
	}
}
