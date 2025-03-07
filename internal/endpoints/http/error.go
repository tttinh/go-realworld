package httpendpoints

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrAccessForbidden = errors.New("access forbidden")
	ErrNotFound        = errors.New("entity not found")
	ErrInternal        = errors.New("internal server error")
	ErrWrongPassword   = errors.New("wrong password")
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

func ok(c *gin.Context, res any) {
	c.JSON(http.StatusOK, res)
}

func error400(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, newErrorRes(err))
}

func error403(c *gin.Context) {
	c.JSON(http.StatusForbidden, newErrorRes(ErrAccessForbidden))
}

func error404(c *gin.Context) {
	c.JSON(http.StatusNotFound, newErrorRes(ErrNotFound))
}

func error500(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, newErrorRes(ErrInternal))
}
