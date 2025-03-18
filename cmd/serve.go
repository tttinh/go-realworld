package cmd

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/tinhtt/go-realworld/internal/adapters"
	"github.com/tinhtt/go-realworld/internal/adapters/postgres"
	"github.com/tinhtt/go-realworld/internal/endpoints"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the API",
	Long:  "Serve the API",
	Run: func(_ *cobra.Command, _ []string) {
		serve()
	},
}

func serve() {
	l.Info("server is starting")

	// Connect database.
	db, err := adapters.ConnectDB(c)
	if err != nil {
		l.Error("fail to connect database", "err", err)
		os.Exit(1)
	}
	defer adapters.CloseDB(db)
	l.Info("connect database successfully")

	// Run migration.
	if c.Migration.Auto {
		err = adapters.Migrate(c)
		if err != nil {
			l.Error("fail to run migration", "err", err)
			os.Exit(1)
		}
		l.Info("run migration successfully")
	}

	// Init repositories, services, handlers.
	users := postgres.NewUsers(db)
	articles := postgres.NewArticles(db)
	server := endpoints.NewHTTPServer(l, c, users, articles)

	// Start HTTP server.
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Error("fail to listen", "err", err)
			os.Exit(1)
		}
	}()

	l.Info("server is running", "port", c.HTTP.Port)

	waitForInterrupt()

	l.Info("server is stopping")

	shutdownGracefully(5*time.Second, func(ctx context.Context) {
		if err := server.Shutdown(ctx); err != nil {
			l.Error("fail to shutdown", "err", err)
			os.Exit(1)
		}
	})

	l.Info("server is exited")
}
