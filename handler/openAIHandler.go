package handler

import (
	"backend/models/requests"
	"backend/models/response"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateTaskName(c *gin.Context) {
	key := "Generate_Task_Name"
	request := &requests.StringRequest{}
	response := &response.StringListResponse{}
	ctx, err := request.Initiate(c, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request, err))
		return
	}

	if err := request.Validate(ctx); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request, err))
		return
	}

	services.GenerateTaskName(ctx, request.Str)

	c.JSON(http.StatusOK, "")
}
