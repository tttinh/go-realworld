package httpport

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type Server struct {
	hs *http.Server
}

func NewServer(
	users domain.UserRepo,
	articles domain.ArticleRepo,
	comments domain.CommentRepo,
) Server {
	handler := gin.Default()
	router := handler.Group("/api")

	ah := articleHandler{articles: articles}
	ah.mount(router)

	ch := commentHandler{articles: articles, comments: comments}
	ch.mount(router)

	uh := userHandler{users: users}
	uh.mount(router)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return Server{hs: s}
}

func (s Server) Run() error {
	return s.hs.ListenAndServe()
}
