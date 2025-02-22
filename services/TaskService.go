package services

import (
	"backend/models/dto"
	"backend/models/entity"
	"backend/models/requests"
	"context"
	"sync"
)

func CreateTask(ctx *context.Context, request *requests.CreateTaskRequest) error {
	if request.Progress == "" {
		request.Progress = "Todo"
	}

	if request.Priority == "" {
		request.Priority = "Low"
	}

	Task := entity.Task{
		Task: dto.Task{
			Title:         request.Title,
			Description:   request.Description,
			Priority:      request.Priority,
			Deadline:      request.Deadline,
			Progress:      request.Progress,
			AssigneeEmail: request.AssigneeEmail,
			CreatedBy:     request.CreatedBy,
		},
	}

	return Task.CreateTask(ctx)
}

func FetchTaskByEmail(ctx *context.Context, email string) ([]*dto.Task, []*dto.Task, error) {
	var TaskAssigned, AssignedTask []*dto.Task

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		TaskAssigned, _ = entity.GetTaskByAssigneEmail(ctx, email)
	}()

	go func() {
		defer wg.Done()
		AssignedTask, _ = entity.GetTaskByCreatedByEmail(ctx, email)
	}()

	wg.Wait()

	return TaskAssigned, AssignedTask, nil
}
