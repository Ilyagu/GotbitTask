package repository

import (
	"context"
	"gotbittask/internal/app/task"
	"gotbittask/internal/app/task/models"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	CreateTaskQuery = `insert into tasks (creator, title, description, executor)
	values ($1,$2,$3,$4)
	returning id, creator, title, description, executor, created, completed;`

	DeleteTaskQuery = "delete from tasks where id=$1;"

	GetTaskQuery = `select id, creator, title, description, executor, created, completed
	from tasks where id=$1;`

	GetAllTasksQuery = `select id, creator, title, description, executor, created, completed
	from tasks order by id;`

	MarkTaskQuery = `update tasks set completed=$2 where id=$1
	returning id, creator, title, description, executor, created, completed;`
)

type PostgreTaskRepo struct {
	Conn *pgxpool.Pool
}

func NewTaskRepo(con *pgxpool.Pool) task.TaskRepository {
	return &PostgreTaskRepo{con}
}

func (tr *PostgreTaskRepo) CreateTask(task models.Task) (models.Task, error) {
	var newTask models.Task
	createdTime := &time.Time{}
	err := tr.Conn.QueryRow(context.Background(), CreateTaskQuery,
		task.Creator, task.Title, task.Description, task.Executor).
		Scan(&newTask.Id, &newTask.Creator, &newTask.Title, &newTask.Description,
			&newTask.Executor, createdTime, &newTask.Completed)
	if err != nil {
		return models.Task{}, err
	}

	newTask.Created = strfmt.DateTime(createdTime.UTC()).String()

	return newTask, nil
}

func (tr *PostgreTaskRepo) DeleteTask(id int64) error {
	_, err := tr.Conn.Exec(context.Background(), DeleteTaskQuery, id)
	return err
}

func (tr *PostgreTaskRepo) GetTask(id int64) (models.Task, error) {
	var existsTask models.Task
	createdTime := &time.Time{}
	err := tr.Conn.QueryRow(context.Background(), GetTaskQuery, id).
		Scan(&existsTask.Id, &existsTask.Creator, &existsTask.Title, &existsTask.Description,
			&existsTask.Executor, createdTime, &existsTask.Completed)
	if err != nil {
		return models.Task{}, err
	}

	existsTask.Created = strfmt.DateTime(createdTime.UTC()).String()

	return existsTask, nil
}

func (tr *PostgreTaskRepo) GetAllTasks() ([]models.Task, error) {
	rows, err := tr.Conn.Query(context.Background(), GetAllTasksQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]models.Task, 0)

	for rows.Next() {
		var task models.Task
		createdTime := &time.Time{}

		err := rows.Scan(&task.Id, &task.Creator, &task.Title,
			&task.Description, &task.Executor, createdTime, &task.Completed)
		if err != nil {
			return nil, err
		}

		task.Created = strfmt.DateTime(createdTime.UTC()).String()
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (tr *PostgreTaskRepo) MarkTask(id int64, mark bool) (models.Task, error) {
	var markedTask models.Task
	createdTime := &time.Time{}

	err := tr.Conn.QueryRow(context.Background(), MarkTaskQuery, id, mark).
		Scan(&markedTask.Id, &markedTask.Creator, &markedTask.Title, &markedTask.Description,
			&markedTask.Executor, createdTime, &markedTask.Completed)
	if err != nil {
		return models.Task{}, err
	}

	markedTask.Created = strfmt.DateTime(createdTime.UTC()).String()

	return markedTask, nil
}
