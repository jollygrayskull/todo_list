package task

import (
	"github.com/jollygrayskull/todo_list/internal/storage"
	"github.com/jollygrayskull/todo_list/internal/storage/models"
	"github.com/jollygrayskull/todo_list/web"
	"html/template"
	"log/slog"
	"net/http"
	"strconv"
)

func RegisterTaskHandlers(mux *http.ServeMux, storage *storage.Storage) {
	mux.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			contentType := r.Header.Get("Content-Type")
			if contentType == "application/x-www-form-urlencoded" {
				createTask(w, r, storage)
			} else {
				slog.Error("Invalid content sent to /task", "content-type", contentType)
				http.Error(w, "Unsupported media type:"+contentType, http.StatusUnsupportedMediaType)
			}
		} else {
			slog.Error("Invalid method used on /task", "method", r.Method)
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

	})

	mux.HandleFunc("/task/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getTask(w, r, storage)
		} else if r.Method == http.MethodDelete {
			deleteTask(w, r, storage)
		} else {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
	})

	mux.HandleFunc("/task/{id}/complete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			completeTask(w, r, storage)
		} else {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
	})
}

func createTask(w http.ResponseWriter, r *http.Request, storage *storage.Storage) {
	taskRepo := storage.TaskRepo

	task, err := models.NewTask(
		r.PostFormValue("title"),
		r.PostFormValue("description"),
		r.PostFormValue("priority"),
		r.PostFormValue("category"))

	if err != nil {
		slog.Error("Invalid task provided to /task", "error", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	_, err = taskRepo.Create(&task)

	if err != nil {
		slog.Error("Failed to create task", "task", task, "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	templ := template.Must(template.ParseFS(web.ViewsFolder, "views/fragments/table_row.html"))
	err = templ.Execute(w, task)
}

func completeTask(w http.ResponseWriter, r *http.Request, storage *storage.Storage) {
	taskRepo := storage.TaskRepo
	idPathVal := r.PathValue("id")

	id, err := strconv.ParseInt(idPathVal, 10, 64)
	if err != nil {
		slog.Error("Path param for id was not a valid integer", "id", idPathVal)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	task, err := taskRepo.Read(int(id))
	if err != nil {
		slog.Error("Failed to complete task, task does not exist", "id", idPathVal)
		http.Error(w, "Task does not exist", http.StatusNotFound)
		return
	}

	task.Completed = true

	err = taskRepo.Update(task)
	if err != nil {
		slog.Error("Failed to complete task, could not update task in store", "id", idPathVal)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	templ := template.Must(template.ParseFS(web.ViewsFolder, "views/fragments/table_row.html"))
	err = templ.Execute(w, task)
}

func getTask(w http.ResponseWriter, r *http.Request, storage *storage.Storage) {
	taskRepo := storage.TaskRepo
	idPathVal := r.PathValue("id")

	id, err := strconv.ParseInt(idPathVal, 10, 64)
	if err != nil {
		slog.Error("Path param for id was not a valid integer", "id", idPathVal)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	task, err := taskRepo.Read(int(id))
	if err != nil {
		slog.Error("Unable to read task with id", "id", id, "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	templ := template.Must(template.ParseFS(web.ViewsFolder, "views/fragments/table_row.html"))
	err = templ.Execute(w, task)
}

func deleteTask(w http.ResponseWriter, r *http.Request, storage *storage.Storage) {
	taskRepo := storage.TaskRepo
	idPathVal := r.PathValue("id")

	id, err := strconv.ParseInt(idPathVal, 10, 64)
	if err != nil {
		slog.Error("Path param for id was not a valid integer", "id", idPathVal)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = taskRepo.Delete(int(id))
	if err != nil {
		slog.Error("Unable to delete task with id", "id", id, "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
