package requests

import (
	"context"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserProfileRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (request *UserProfileRequest) Initiate(c *gin.Context, key string) (*context.Context, error) {
	_ctx, _ := c.Get("context")
	ctx := _ctx.(context.Context)

	if err := c.BindJSON(&request); err != nil {
		return &ctx, err
	}

	return &ctx, nil
}

func (request *UserProfileRequest) Validate(ctx *context.Context) error {
	request.Email = strings.TrimSpace(request.Email)
	request.Name = strings.TrimSpace(request.Name)

	if len(request.Name) == 0 {
		return errors.New("name cann't be empty")
	}

	if len(request.Email) == 0 {
		return errors.New("enter a valid email")
	}

	if len(request.Password) == 0 {
		return errors.New("enter a valid password")
	}

	return nil
}
