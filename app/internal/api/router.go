package api

import (
	"net/http"
	api "taskmanager/app/internal/api/tasks"

	"github.com/go-chi/chi/v5"
)

func Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) { api.getAllTasks(repo, w) })
		r.Post("/", api.CreateTask)
		r.Get("/{id}", api.GetTaskByID)
	})

	return r
}
