package usecase

import (
	"gotbittask/internal/app/task"
	"gotbittask/internal/app/task/models"
)

type TaskUsecase struct {
	taskRepo task.TaskRepository
}

func NewTaskUsecase(tr task.TaskRepository) task.TaskUsecase {
	return &TaskUsecase{
		taskRepo: tr,
	}
}

func (tu *TaskUsecase) CreateTask(task models.Task) (models.Task, error) {
	return tu.taskRepo.CreateTask(task)
}

func (tu *TaskUsecase) DeleteTask(id int64) error {
	_, err := tu.taskRepo.GetTask(id)
	if err != nil {
		return err
	}
	return tu.taskRepo.DeleteTask(id)
}

func (tu *TaskUsecase) GetAllTasks() ([]models.Task, error) {
	return tu.taskRepo.GetAllTasks()
}

func (tu *TaskUsecase) MarkTask(id int64, mark bool) (models.Task, error) {
	_, err := tu.taskRepo.GetTask(id)
	if err != nil {
		return models.Task{}, err
	}
	return tu.taskRepo.MarkTask(id, mark)
}
