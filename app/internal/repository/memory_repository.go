package repository

import (
	"errors"
	"taskmanager/app/internal/domain/task"
)

type MemoryTaskRepository struct {
	memory []task.Task
}

func NewMemoryTaskRepository() *MemoryTaskRepository {
	return &MemoryTaskRepository{memory: []task.Task{}}
}

func (r *MemoryTaskRepository) AddTask(task task.Task) error {
	r.memory = append(r.memory, task)
	return nil
}

func (r *MemoryTaskRepository) GetTaskById(id int) (task.Task, error) {
	for _, task := range r.memory {
		if task.Id == id {
			return task, nil
		}
	}
	return task.Task{}, errors.New("task not found")
}

func (r *MemoryTaskRepository) GetAllTasks() ([]task.Task, error) {
	return r.memory, nil
}
