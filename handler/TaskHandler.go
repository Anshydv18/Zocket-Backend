package handler

import (
	"backend/models/requests"
	"backend/models/response"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	key := "Create_Task"
	request := &requests.CreateTaskRequest{}
	response := &response.BaseResponse{}

	ctx, err := request.Initiate(c, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request))
		return
	}

	if err := request.Validate(ctx); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request))
		return
	}

	if err := services.CreateTask(ctx, request); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request))
		return
	}

	c.JSON(http.StatusOK, response.Success(ctx))
}

func FetchTasks(c *gin.Context) {
	key := "fetch_tasks"
	request := &requests.StringRequest{}
	response := &response.TaskListResponse{}

	ctx, err := request.Initiate(c, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request))
		return
	}

	if err := request.Validate(ctx); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request))
		return
	}

	TaskAssigned, AssignedTask, err := services.FetchTaskByEmail(ctx, request.Str)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request))
		return
	}

	c.JSON(http.StatusOK, response.Success(ctx, TaskAssigned, AssignedTask))
}
