package container

import (
	"taskmanager/app/internal/config"
	"taskmanager/app/internal/repository"
)

type AppDependencies struct {
	AppConfig      *config.Config
	TaskRepository repository.BaseTaskRepository
}

func InitContainer(cfg *config.Config) AppDependencies {
	container := AppDependencies{
		AppConfig:      cfg,
		TaskRepository: repository.NewPostgresTaskRepository(cfg.GetDatabaseDSN()),
	}
	return container
}
