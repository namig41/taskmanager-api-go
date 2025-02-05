package repository

import "taskmanager/app/internal/domain/task"

type BaseTaskRepository interface {
	AddTask(task.Task) error
	GetTaskById(int) (task.Task, error)
	GetAllTasks() ([]task.Task, error)
}
