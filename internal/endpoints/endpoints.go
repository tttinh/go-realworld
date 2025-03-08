package endpoints

import (
	"net/http"
	"time"

	"github.com/tinhtt/go-realworld/internal/domain"
	httpendpoints "github.com/tinhtt/go-realworld/internal/endpoints/http"
)

func NewHTTPServer(
	users domain.UserRepo,
	articles domain.ArticleRepo,
) *http.Server {
	return &http.Server{
		Addr: ":8080",
		Handler: httpendpoints.NewHandler(
			users,
			articles,
		),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
