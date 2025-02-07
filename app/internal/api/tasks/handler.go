package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"taskmanager/app/internal/container"
	"taskmanager/app/internal/domain/task"

	"github.com/go-chi/chi/v5"
)

var (
	mutex sync.Mutex
)

func GetAllTasks(app *container.AppDependencies, w http.ResponseWriter, r *http.Request) {
	mutex.Lock()

	taskList, err := app.TaskRepository.GetAllTasks()

	if err != nil {
		fmt.Print("Get All Tasks Error")
		return
	}

	mutex.Unlock()

	jsonResponse(w, http.StatusOK, taskList)
}

func CreateTask(app *container.AppDependencies, w http.ResponseWriter, r *http.Request) {
	var newTask task.Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	mutex.Lock()

	err := app.TaskRepository.AddTask(newTask)

	if err != nil {
		fmt.Print("Create Task Error")
	}

	mutex.Unlock()

	jsonResponse(w, http.StatusCreated, newTask)
}

func GetTaskByID(app *container.AppDependencies, w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mutex.Lock()

	task, err := app.TaskRepository.GetTaskById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mutex.Unlock()

	jsonResponse(w, http.StatusOK, task)
}

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func parseIDParam(r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "id")
	var id int
	_, err := fmt.Sscanf(idStr, "%d", &id)
	if err != nil {
		return 0, fmt.Errorf("Invalid task ID")
	}
	return id, nil
}
