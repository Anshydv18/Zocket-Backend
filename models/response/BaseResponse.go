package response

import (
	"context"
	"net/http"
)

type BaseResponse struct {
	Status     bool        `json:"status"`
	StatusCode int         `json:"status_code"`
	Error      string      `json:"error" omitempty`
	Data       interface{} `json:"request"`
}

func (response *BaseResponse) Success(ctx *context.Context) *BaseResponse {
	response.Status = true
	response.StatusCode = http.StatusOK
	return response
}

func (response *BaseResponse) Fail(ctx *context.Context, request interface{}, err error) *BaseResponse {
	response.Status = false
	response.Data = request
	response.Error = err.Error()
	return response
}
