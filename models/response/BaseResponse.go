package response

import (
	"context"
	"net/http"
)

type BaseResponse struct {
	Status     bool        `json:"status"`
	StatusCode int         `json:"status_code"`
	Error      error       `json:"error" omitempty`
	Data       interface{} `json:"request"`
}

func (response *BaseResponse) Success(ctx *context.Context) *BaseResponse {
	response.Status = true
	response.StatusCode = http.StatusOK
	return response
}

func (response *BaseResponse) Fail(ctx *context.Context, request interface{}) *BaseResponse {
	response.Status = true
	response.Data = request
	return response
}
