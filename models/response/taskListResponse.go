package response

import (
	"backend/models/dto"
	"context"
)

type TaskListResponse struct {
	BaseResponse
	TaskAssigned []*dto.Task `json:"tasks_assigned"`
}

func (res *TaskListResponse) Success(ctx *context.Context, data []*dto.Task) *TaskListResponse {
	res.StatusCode = 200
	res.Status = true
	res.TaskAssigned = data
	return res
}
