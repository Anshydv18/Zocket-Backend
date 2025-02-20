package response

import (
	"backend/models/dto"
	"context"
)

type UserResponse struct {
	BaseResponse
	UserData *dto.UserDto
}

func (response *UserResponse) Success(ctx *context.Context, data *dto.UserDto) *UserResponse {
	response.Status = true
	response.StatusCode = 200
	response.UserData = data
	return response
}
