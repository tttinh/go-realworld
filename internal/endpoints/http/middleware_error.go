package httpendpoints

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorRes struct {
	Errors struct {
		Body []string `json:"body"`
	} `json:"errors"`
}

func newErrorRes(args ...error) errorRes {
	var res errorRes
	for _, err := range args {
		res.Errors.Body = append(res.Errors.Body, err.Error())
	}
	return res
}

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			if c.Writer.Status() == http.StatusInternalServerError {
				c.JSON(-1, newErrorRes(errors.New("internal server error")))
				return
			}

			c.JSON(-1, newErrorRes(err))
		}
	}
}
