package task

import "gotbittask/internal/app/task/models"

type TaskRepository interface {
	CreateTask(task models.Task) (models.Task, error)
	DeleteTask(id int64) error
	GetAllTasks() ([]models.Task, error)
	MarkTask(id int64, mark bool) (models.Task, error)
	GetTask(id int64) (models.Task, error)
}

type TaskUsecase interface {
	CreateTask(task models.Task) (models.Task, error)
	DeleteTask(id int64) error
	GetAllTasks() ([]models.Task, error)
	MarkTask(id int64, mark bool) (models.Task, error)
}
