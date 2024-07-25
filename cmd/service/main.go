package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/albertyla/connectisend/internal/service"
	"github.com/albertyla/connectisend/internal/service/config"
	"github.com/albertyla/connectisend/internal/util"
	"golang.org/x/sync/errgroup"
)

const (
	readTimeout     = 10 * time.Second
	shutdownTimeout = 10 * time.Second
)

func main() {
	ctx := context.Background()
	if err := run(ctx, os.LookupEnv, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, lookupenv func(string) (string, bool), w io.Writer) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	conf, err := config.NewConfig(ctx, lookupenv)
	if err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}

	logger := util.NewLogger(w, conf)
	srv := service.NewServer(
		logger,
		conf,
	)

	httpServer := &http.Server{
		ReadHeaderTimeout: readTimeout,
		Addr:              net.JoinHostPort(conf.Host, conf.Port),
		Handler:           srv,
	}

	var eg errgroup.Group
	eg.Go(func() error {
		logger.InfoContext(ctx, "listening", "httpserver.Addr", httpServer.Addr)
		if serverErr := httpServer.ListenAndServe(); serverErr != nil && serverErr != http.ErrServerClosed {
			return fmt.Errorf("error listening and serving: %w", serverErr)
		}
		return nil
	})

	eg.Go(func() error {
		<-ctx.Done()
		// make a new context for the Shutdown
		shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancelShutdown()
		if shutdownErr := httpServer.Shutdown(shutdownCtx); shutdownErr != nil {
			return fmt.Errorf("error shutting down http server: %w", shutdownErr)
		}
		return nil
	})

	return eg.Wait()
}
