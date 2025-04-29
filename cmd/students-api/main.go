package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/thenaveensharma/students-api/internal/config"
	"github.com/thenaveensharma/students-api/internal/http/handlers/student"
	"github.com/thenaveensharma/students-api/internal/storage/sqlite"
)

func main() {
	//load config
	cfg := config.MustLoad()

	//databse initialize
	storage, err := sqlite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}
	slog.Info("storage initialized", slog.String("env", cfg.Env))
	//server

	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	slog.Info("server started", slog.String("address", cfg.Address))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			slog.Error("failed to start server", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	sig := <-done
	slog.Info("shutting down the server", slog.String("signal", sig.String()))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	} else {
		slog.Info("server shutdown successfully")
	}

}
