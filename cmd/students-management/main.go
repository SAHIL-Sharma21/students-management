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

	"github.com/SAHIL-Sharma21/students-management/pkg/config"
)

func main() {
	//load config
	cfg := config.MustLoadConfig()

	//database setup

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to students management"))
	})
	//server setup
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	slog.Info("Server started", slog.String("address", cfg.Address))

	//NOTE:production mei gracfuly shutdown krna hai server ko agr interrupt signal aata hai toh
	//running the server in goroutine

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("failed to start server: ", err)
		}
	}()

	<-done

	//logic for server stop
	//graceful shutdown in production
	slog.Info("shutting down the server.....")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server....", slog.String("error", err.Error()))
	}

	slog.Info("server stopped successfully!")
}
