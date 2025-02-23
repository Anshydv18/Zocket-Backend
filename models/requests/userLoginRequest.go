package requests

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
)

type UserLoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (request *UserLoginRequest) Initiate(c *gin.Context, key string) (*context.Context, error) {
	_ctx, _ := c.Get("context")
	ctx := _ctx.(context.Context)

	if err := c.ShouldBindJSON(&request); err != nil {
		return &ctx, err
	}

	return &ctx, nil
}

func (request *UserLoginRequest) Validate(ctx *context.Context) error {
	// request.Email = strings.TrimSpace(request.Email)
	if len(request.Email) == 0 {
		return errors.New("enter a valid email")
	}

	if len(request.Password) == 0 {
		return errors.New("enter a valid password")
	}

	return nil
}
