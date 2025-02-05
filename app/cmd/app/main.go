package main

import (
	"fmt"
	"log"
	"taskmanager/app/internal/config"
	"taskmanager/app/internal/domain/task"
	"taskmanager/app/internal/repository"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(cfg)

	// Создаем репозиторий
	repo := repository.NewMemoryTaskRepository()

	// Тестируем добавление задачи
	task1 := task.Task{Id: 1, Name: "Test Task 1"}

	repo.AddTask(task1)

	_, _ = repo, task1
}
