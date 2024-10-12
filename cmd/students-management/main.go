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
	"github.com/SAHIL-Sharma21/students-management/pkg/http/handlers/student"
	"github.com/SAHIL-Sharma21/students-management/pkg/storage/sqlite"
)

func main() {
	//load config
	cfg := config.MustLoadConfig()

	//database setup
	storage, err := sqlite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.1"))
	//setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/v1/students", student.New(storage))
	router.HandleFunc("GET /api/v1/students/{id}", student.GetById(storage))

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

	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server....", slog.String("error", err.Error()))
	}

	slog.Info("server stopped successfully!")
}
