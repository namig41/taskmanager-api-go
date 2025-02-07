package main

import (
	"log"
	"net/http"
	"taskmanager/app/internal/api"
	"taskmanager/app/internal/config"
	"taskmanager/app/internal/container"
	"taskmanager/app/internal/domain/task"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	deps := container.InitContainer(cfg)

	if err != nil {
		log.Fatal(err)
	}

	task1 := task.Task{Id: 1, Name: "Test Task 1"}

	deps.TaskRepository.AddTask(task1)

	r := api.Routes()

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
