package repository

import (
	"database/sql"
	"errors"
	"taskmanager/app/internal/domain/task"

	_ "github.com/lib/pq"
)

type PostgresTaskRepository struct {
	db *sql.DB
}

func NewPostgresTaskRepository(dsn string) *PostgresTaskRepository {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil
	}

	if err := db.Ping(); err != nil {
		return nil
	}

	return &PostgresTaskRepository{db: db}
}

func (r *PostgresTaskRepository) AddTask(t task.Task) error {
	query := "INSERT INTO tasks (id, name) VALUES ($1, $2)"
	_, err := r.db.Exec(query, t.Id, t.Name)
	return err
}

func (r *PostgresTaskRepository) GetTaskById(id int) (task.Task, error) {
	query := "SELECT id, name FROM tasks WHERE id = $1"
	var t task.Task
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&t.Id, &t.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return task.Task{}, errors.New("task not found")
		}
		return task.Task{}, err
	}
	return t, nil
}

func (r *PostgresTaskRepository) GetAllTasks() ([]task.Task, error) {
	query := "SELECT id, name FROM tasks"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []task.Task{}
	for rows.Next() {
		var t task.Task
		if err := rows.Scan(&t.Id, &t.Name); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *PostgresTaskRepository) Close() error {
	return r.db.Close()
}
