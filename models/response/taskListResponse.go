package response

import (
	"backend/models/dto"
	"context"
)

type TaskListResponse struct {
	BaseResponse
	TaskAssigned []*dto.Task `json:"tasks_assigned"`
	AssignedTask []*dto.Task `json:"assigned_tasks"`
}

func (res *TaskListResponse) Success(ctx *context.Context, data []*dto.Task, data2 []*dto.Task) *TaskListResponse {
	res.StatusCode = 200
	res.Status = true
	res.TaskAssigned = data
	res.AssignedTask = data2
	return res
}
