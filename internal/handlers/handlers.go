package handlers

import (
	"github.com/jollygrayskull/todo_list/internal/handlers/root"
	"github.com/jollygrayskull/todo_list/internal/handlers/task"
	"github.com/jollygrayskull/todo_list/internal/storage"
	"net/http"
)

func RegisterHandlers(mux *http.ServeMux, storage *storage.Storage) {
	root.RegisterRootHandlers(mux, storage)
	task.RegisterTaskHandlers(mux, storage)
}
