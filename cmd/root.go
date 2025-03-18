package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/tinhtt/go-realworld/internal/config"
	"github.com/tinhtt/go-realworld/internal/pkg"
)

var l *slog.Logger
var c *config.Config
var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A command line program",
	Long:  "This application implements the Conduit API",
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initLogger, loadConfig)
	rootCmd.AddCommand(serveCmd)
}

func initLogger() {
	m := config.Mode(os.Getenv("CONDUIT_MODE"))
	l = pkg.NewLogger(m)
}

func loadConfig() {
	var err error
	c, err = config.Load()
	if err != nil {
		l.Error("fail to load config", "err", err)
		os.Exit(1)
	}

	// In production, do not print out the secrets.
	l.Info("load config successfully", "cfg", c)
}

func waitForInterrupt() {
	// Wait for interrupt signal to gracefully shutdown the application
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func shutdownGracefully(timeout time.Duration, shutdownFunc func(context.Context)) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	go func() {
		shutdownFunc(ctx)

		// Task completed, cancel the context.
		cancel()
	}()

	<-ctx.Done()

	if ctx.Err() == context.DeadlineExceeded {
		l.Warn(fmt.Sprintf("fail to shutdown within %d miliseconds", timeout.Milliseconds()))
		return
	}

	l.Info("shutdown gracefully")
}
