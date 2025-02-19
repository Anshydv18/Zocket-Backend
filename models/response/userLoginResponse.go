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
	response.UserData = data
	return response
}
