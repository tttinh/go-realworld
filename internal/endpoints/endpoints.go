package endpoints

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/tinhtt/go-realworld/internal/config"
	"github.com/tinhtt/go-realworld/internal/domain"
	httpendpoints "github.com/tinhtt/go-realworld/internal/endpoints/http"
)

func NewHTTPServer(
	log *slog.Logger,
	cfg config.Config,
	users domain.UserRepo,
	articles domain.ArticleRepo,
) *http.Server {
	return &http.Server{
		Addr: cfg.HTTP.Port,
		Handler: httpendpoints.NewHandler(
			log,
			cfg,
			users,
			articles,
		),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
