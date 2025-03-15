package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/tinhtt/go-realworld/internal/adapters"
	"github.com/tinhtt/go-realworld/internal/adapters/postgres"
	"github.com/tinhtt/go-realworld/internal/endpoints"
	"github.com/tinhtt/go-realworld/internal/pkg"
)

func main() {
	log := pkg.NewLogger(os.Getenv("MODE"))

	db := adapters.ConnectDB()
	defer adapters.CloseDB(db)

	users := postgres.NewUsers(db)
	articles := postgres.NewArticles(db)

	server := endpoints.NewHTTPServer(log, users, articles)

	log.Info("server starting")
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("fail to listen", "err", err)
			os.Exit(1)
		}
	}()

	log.Info("server running")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("server shutting down")

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
	log.Info("server exiting")
}
