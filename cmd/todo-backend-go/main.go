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

	"github.com/gauravst/todo-backend-go/internal/api/handlers"
	"github.com/gauravst/todo-backend-go/internal/config"
	"github.com/gauravst/todo-backend-go/internal/database"
	"github.com/gauravst/todo-backend-go/internal/repositories"
	"github.com/gauravst/todo-backend-go/internal/services"
)

func main() {
	// load config
	cfg := config.ConfigMustLoad()

	// database setup
	database.InitDB(cfg.DatabaseUri)
	defer database.CloseDB()

	//setup router
	router := http.NewServeMux()

	userRepo := repositories.NewTaskRepository(database.DB)
	userService := services.NewTaskService(userRepo)

	router.HandleFunc("GET /api/task", handlers.GetTask(userService))
	router.HandleFunc("GET /api/task/{id}", handlers.GetTask(userService))
	router.HandleFunc("POST /api/task", handlers.CreateTask(userService))
	router.HandleFunc("PUT /api/task/{id}", handlers.UpdateTask(userService))
	router.HandleFunc("DELETE /api/task/{id}", handlers.DeleteTask(userService))

	// setup server
	server := &http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	slog.Info("server started", slog.String("address", cfg.Address))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("faild to Shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server Shutdown successfully")
}
