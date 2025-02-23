package requests

import (
	"context"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

type StringRequest struct {
	Str string `json:"string"`
}

func (req *StringRequest) Initiate(c *gin.Context, key string) (*context.Context, error) {
	_ctx, _ := c.Get("context")
	ctx := _ctx.(context.Context)

	if err := c.ShouldBindJSON(&req); err != nil {
		return &ctx, err
	}

	return &ctx, nil
}

func (req *StringRequest) Validate(ctx *context.Context) error {
	req.Str = strings.TrimSpace(req.Str)
	if len(req.Str) == 0 {
		return errors.New("enter a valid input")
	}

	return nil
}
