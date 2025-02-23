package handler

import (
	"backend/models/requests"
	"backend/models/response"
	"backend/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	key := "Create_Task"
	request := &requests.CreateTaskRequest{}
	response := &response.BaseResponse{}

	ctx, err := request.Initiate(c, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request, err))
		return
	}

	if err := request.Validate(ctx); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request, err))
		return
	}

	if err := services.CreateTask(ctx, request); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request, err))
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
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request, err))
		return
	}

	if err := request.Validate(ctx); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request, err))
		return
	}

	TaskAssigned, AssignedTask, err := services.FetchTaskByEmail(ctx, request.Str)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request, err))
		return
	}

	c.JSON(http.StatusOK, response.Success(ctx, TaskAssigned, AssignedTask))
}

func UpdateTask(c *gin.Context) {
	key := "update_Task"
	request := &requests.CreateTaskRequest{}
	response := &response.BaseResponse{}

	ctx, err := request.Initiate(c, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request, err))
		return
	}

	if err := request.Validate(ctx); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request, err))
		return
	}

	if err := request.ValidateId(ctx); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request, err))
		return
	}

	if err := services.UpdateTask(ctx, request); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request, err))
		return
	}

	NotifyUser(request.CreatedBy, fmt.Sprintf("Task %s has been changed", request.Title))

	c.JSON(http.StatusOK, response.Success(ctx))
}
