package requests

import (
	"backend/models/dto"
	"context"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

type CreateTaskRequest struct {
	dto.Task
}

func (req *CreateTaskRequest) Initiate(c *gin.Context, key string) (*context.Context, error) {
	_ctx, _ := c.Get("context")
	ctx := _ctx.(context.Context)
	if err := c.BindJSON(req); err != nil {
		return &ctx, err
	}

	return &ctx, nil
}

func (req *CreateTaskRequest) Validate(ctx *context.Context) error {
	req.Title = strings.TrimSpace(req.Title)
	req.AssigneeEmail = strings.TrimSpace(req.AssigneeEmail)

	if len(req.Title) == 0 {
		return errors.New("please enter title")
	}

	if len(req.AssigneeEmail) == 0 {
		return errors.New("please enter assignee email")
	}

	if len(req.CreatedBy) == 0 {
		return errors.New("please Login Again")
	}
	return nil
}
