package response

import "context"

type StringListResponse struct {
	BaseResponse
	List []string `json:"list"`
}

func (res *StringListResponse) Success(ctx *context.Context, list []string) *StringListResponse {
	res.StatusCode = 200
	res.Status = true
	res.List = list
	return res
}
