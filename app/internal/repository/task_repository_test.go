package repository

import (
	"taskmanager/app/internal/domain/task"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryTaskRepository(t *testing.T) {
	repo := NewMemoryTaskRepository()

	task1 := task.Task{Id: 1, Name: "Test Task 1"}
	err := repo.AddTask(task1)
	assert.Nil(t, err, "AddTask should not return error")

	t.Run("GetTaskById", func(t *testing.T) {
		t.Run("task exists", func(t *testing.T) {
			foundTask, err := repo.GetTaskById(1)
			assert.Nil(t, err, "GetTaskById should not return error")
			assert.Equal(t, task1, foundTask, "Found task should match the added task")
		})

		t.Run("task not found", func(t *testing.T) {
			_, err := repo.GetTaskById(999)
			assert.NotNil(t, err, "GetTaskById should return error when task not found")
		})
	})

	// Тестируем получение всех задач
	t.Run("GetAllTasks", func(t *testing.T) {
		tasks, err := repo.GetAllTasks()
		assert.Nil(t, err, "GetAllTasks should not return error")
		assert.Len(t, tasks, 1, "Should return 1 task")
		assert.Equal(t, task1, tasks[0], "Returned task should match the added task")
	})

	// Тестируем добавление второй задачи
	task2 := task.Task{Id: 2, Name: "Test Task 2"}
	err = repo.AddTask(task2)
	assert.Nil(t, err, "AddTask should not return error")

	t.Run("GetAllTasks after adding second task", func(t *testing.T) {
		tasks, err := repo.GetAllTasks()
		assert.Nil(t, err, "GetAllTasks should not return error")
		assert.Len(t, tasks, 2, "Should return 2 tasks")
		assert.Equal(t, task2, tasks[1], "Second task should match the added task")
	})
}
