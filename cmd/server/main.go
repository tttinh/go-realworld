package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/tinhtt/go-realworld/internal/adapters"
	"github.com/tinhtt/go-realworld/internal/adapters/postgres"
	"github.com/tinhtt/go-realworld/internal/endpoints"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	db := adapters.ConnectDB()
	defer adapters.CloseDB(db)

	users := postgres.NewUsers(db)
	articles := postgres.NewArticles(db)
	server := endpoints.NewHTTPServer(users, articles)

	log.Info().Msg("Start Server ...")
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("listen: %s\n", err)
		}
	}()

	log.Info().Msg("Server running")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutdown Server ...")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server Shutdown:", err)
	// }
	// // catching ctx.Done(). timeout of 5 seconds.
	// select {
	// case <-ctx.Done():
	// 	log.Println("timeout of 5 seconds.")
	// }
	log.Info().Msg("Server exiting")
}
