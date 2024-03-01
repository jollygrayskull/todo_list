package root

import (
	"github.com/jollygrayskull/todo_list/internal/storage"
	"github.com/jollygrayskull/todo_list/internal/storage/models"
	"github.com/jollygrayskull/todo_list/web"
	"html/template"
	"log"
	"log/slog"
	"net/http"
)

func RegisterRootHandlers(mux *http.ServeMux, storage *storage.Storage) {
	serveStaticData(mux)
	registerRootHandler(mux, storage)
}

func registerRootHandler(mux *http.ServeMux, storage *storage.Storage) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			renderIndex(w, storage)
		} else {
			slog.Error("Invalid method used on /", "method", r.Method)
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
	})
}

func renderIndex(w http.ResponseWriter, storage *storage.Storage) {
	taskRepo := storage.TaskRepo

	all, err := taskRepo.ReadAllSorted(sortByPriorityThenIndex)
	if err != nil {
		slog.Error("Failed to read tasks from store", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	templ := template.Must(template.ParseFS(web.ViewsFolder, "views/index.html"))
	data := map[string][]*models.Task{"Tasks": all}
	err = templ.Execute(w, data)
	if err != nil {
		log.Fatal("Failed to parse index file", err)
	}
}

func sortByPriorityThenIndex(left *models.Task, right *models.Task) bool {
	if left.Priority == right.Priority {
		return left.Id < right.Id
	} else {
		return left.Priority < right.Priority
	}
}

func serveStaticData(mux *http.ServeMux) {
	mux.Handle("/static/", http.FileServer(http.FS(web.StaticFolder)))
}
