package services

import (
	"backend/models/dto"
	"backend/models/entity"
	"backend/models/requests"
	"context"
)

func CreateTask(ctx *context.Context, request *requests.CreateTaskRequest) error {
	if request.Progess == "" {
		request.Progess = "Todo"
	}

	if request.Priority == "" {
		request.Priority = "Low"
	}

	Task := entity.Task{
		Title:         request.Title,
		Description:   request.Description,
		Priority:      request.Priority,
		Deadline:      request.Deadline,
		Progess:       request.Progess,
		AssigneeEmail: request.AssigneeEmail,
		CreatedBy:     request.CreatedBy,
	}

	return Task.CreateTask(ctx)
}

func FetchTaskByEmail(ctx *context.Context, email string) ([]*dto.Task, error) {
	return entity.GetTaskByAssigneEmail(ctx, email)
}
