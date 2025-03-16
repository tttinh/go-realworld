package httpendpoints

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/pkg"
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
		defer func() {
			if err := recover(); err != nil {
				e := pkg.NewError(fmt.Errorf("panic: %v", err))
				c.AbortWithError(http.StatusInternalServerError, e)
			}
		}()
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
