package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

func NewHttpHandler(
	users domain.UserRepo,
	articles domain.ArticleRepo,
	comments domain.CommentRepo,
) http.Handler {
	handler := gin.Default()
	router := handler.Group("/api")

	ah := articleHandler{articles: articles}
	ah.mount(router)

	ch := commentHandler{articles: articles, comments: comments}
	ch.mount(router)

	uh := userHandler{users: users}
	uh.mount(router)

	return handler
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
