package delivery

import (
	"encoding/json"
	"gotbittask/internal/app/middlware"
	"gotbittask/internal/app/task"
	"gotbittask/internal/app/task/models"
	"gotbittask/internal/pkg/responses"
	"log"
	"net/http"
	"strconv"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type TaskHandler struct {
	taskUsecase task.TaskUsecase
}

func NewTaskHandler(router *router.Router, tu task.TaskUsecase) {
	taskHandler := &TaskHandler{
		taskUsecase: tu,
	}

	router.POST("/api/v1/task",
		middlware.ReponseMiddlwareAndLogger(taskHandler.CreateTaskHandler))
	router.GET("/api/v1/task/all",
		middlware.ReponseMiddlwareAndLogger(taskHandler.GetAllTasksHandler))
	router.DELETE("/api/v1/task/{id}",
		middlware.ReponseMiddlwareAndLogger(taskHandler.DeleteTaskHandler))
	router.PUT("/api/v1/task/mark",
		middlware.ReponseMiddlwareAndLogger(taskHandler.MarkTaskHandler))
}

// @Summary      CreateTask
// @Description  creating task
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        input body models.Task true "account info"
// @Success      200  {object}  models.Task
// @Failure      400  {object}  responses.Response
// @Failure      404  {object}  responses.Response
// @Failure      500  {object}  responses.Response
// @Router       /task [post]
func (th *TaskHandler) CreateTaskHandler(ctx *fasthttp.RequestCtx) {
	newTask := models.Task{}
	err := json.Unmarshal(ctx.PostBody(), &newTask)
	if err != nil {
		log.Println(err)
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}
	task, err := th.taskUsecase.CreateTask(newTask)
	if err != nil {
		responses.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
	}
	taskBody, err := json.Marshal(task)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		return
	}
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(taskBody)
}

// @Summary      GetAllTasks
// @Description  getting all tasks
// @Tags         task
// @Produce      json
// @Success      200  {object}  models.Task
// @Failure      404  {object}  responses.Response
// @Failure      500  {object}  responses.Response
// @Router       /task/all [get]
func (th *TaskHandler) GetAllTasksHandler(ctx *fasthttp.RequestCtx) {
	tasks, err := th.taskUsecase.GetAllTasks()
	if err != nil {
		responses.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
	}
	tasksBody, err := json.Marshal(tasks)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		return
	}
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(tasksBody)
}

// @Summary      DeleteTask
// @Description  deleting task by id
// @Tags         task
// @Produce      json
// @Success      200  {object}  models.Task
// @Failure      404  {object}  responses.Response
// @Failure      500  {object}  responses.Response
// @Router       /task/{id} [delete]
func (th *TaskHandler) DeleteTaskHandler(ctx *fasthttp.RequestCtx) {
	idStr := ctx.UserValue("id").(string)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		responses.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	err = th.taskUsecase.DeleteTask(int64(id))
	if err != nil {
		responses.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
	}
	ctx.SetStatusCode(http.StatusOK)
}

// @Summary      MarkTas
// @Description  mark task to completed or not completed by id
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        input body models.Task true "account info"
// @Success      200  {object}  models.Task
// @Failure      400  {object}  responses.Response
// @Failure      404  {object}  responses.Response
// @Failure      500  {object}  responses.Response
// @Router       /task/mark [put]
func (th *TaskHandler) MarkTaskHandler(ctx *fasthttp.RequestCtx) {
	taskToMark := models.Task{}
	err := json.Unmarshal(ctx.PostBody(), &taskToMark)
	if err != nil {
		log.Println(err)
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}
	markedTask, err := th.taskUsecase.MarkTask(taskToMark.Id, taskToMark.Completed)
	if err != nil {
		responses.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
	}
	markedTaskBody, err := json.Marshal(markedTask)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		return
	}
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(markedTaskBody)
}
