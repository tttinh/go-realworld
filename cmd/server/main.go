package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/tinhtt/go-realworld/internal/infra"
	pgrepo "github.com/tinhtt/go-realworld/internal/infra/postgres/repo"
	"github.com/tinhtt/go-realworld/internal/port"
)

func main() {
	db := infra.ConnectDB()
	defer infra.CloseDB(db)

	users := pgrepo.NewUsers(db)
	articles := pgrepo.NewArticles(db)
	comments := pgrepo.NewComments(db)
	server := port.NewHTTPServer(users, articles, comments)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

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
	log.Println("Server exiting")
}
