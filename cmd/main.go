package main

import (
	"fmt"
	"github.com/jollygrayskull/todo_list/internal/config"
	"github.com/jollygrayskull/todo_list/internal/handlers"
	"github.com/jollygrayskull/todo_list/internal/storage"
	"github.com/jollygrayskull/todo_list/internal/storage/models"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	conf, err := config.ReadConfig()
	if err != nil {
		slog.Warn("Failed to read config, loading defaults", "error", err)
	}

	serverConf := conf.ServerConfig
	store := storage.NewIMStorage()
	addr := fmt.Sprintf("%s:%d", serverConf.BindAddress, serverConf.BindPort)
	mux := http.NewServeMux()

	handlers.RegisterHandlers(mux, store)

	createExampleData(store)
	slog.Info("Running server on " + addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func createExampleData(store *storage.Storage) {
	taskRepo := store.TaskRepo

	_, _ = taskRepo.Create(&models.Task{
		Title:       "Test task 1",
		Description: "Test description 1",
		Priority:    models.High,
		Category:    models.Personal,
	})
	_, _ = taskRepo.Create(&models.Task{
		Title:       "Test task 2",
		Description: "Test description 2",
		Priority:    models.Medium,
		Category:    models.Work,
	})
	_, _ = taskRepo.Create(&models.Task{
		Title:       "Test task 3",
		Description: "Test description 3",
		Priority:    models.Low,
		Category:    models.Household,
	})
}
