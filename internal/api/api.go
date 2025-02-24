package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/repo"
)

func NewHttpHandler(
	users repo.Users,
	articles repo.Articles,
	comments repo.Comments,
) http.Handler {
	router := gin.Default()

	ah := articleHandler{articles: articles}
	ah.mount(router)

	ch := CommentHandler{articles: articles, comments: comments}
	ch.mount(router)

	uh := UserHandler{users: users}
	uh.mount(router)

	return router
}

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
