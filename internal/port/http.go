package port

import (
	"net/http"
	"time"

	"github.com/tinhtt/go-realworld/internal/domain"
	httpport "github.com/tinhtt/go-realworld/internal/port/http"
)

func NewHTTPServer(
	users domain.UserRepo,
	articles domain.ArticleRepo,
	comments domain.CommentRepo,
) *http.Server {
	return &http.Server{
		Addr: ":8080",
		Handler: httpport.NewHandler(
			users,
			articles,
			comments,
		),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
