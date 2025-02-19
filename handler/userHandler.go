package handler

import (
	"backend/models/requests"
	"backend/models/response"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	key := "User_Login"
	request := &requests.UserLoginRequest{}
	response := &response.UserResponse{}

	ctx, err := request.Initiate(c, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request))
		return
	}

	if err := request.Validate(ctx); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request))
		return
	}

	data, err := services.UserLogin(ctx, request)
	if err != nil {
		c.JSON(http.StatusOK, response.Fail(ctx, request))
	}

	c.JSON(http.StatusOK, response.Success(ctx, data))
}

func CreateUserProfile(c *gin.Context) {
	key := "Create_User_Profile"
	request := &requests.UserProfileRequest{}
	response := &response.BaseResponse{}
	ctx, err := request.Initiate(c, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request))
		return
	}

	if err := request.Validate(ctx); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request))
		return
	}

	if err := services.CreateUserProfile(ctx, request); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request))
		return
	}

	c.JSON(http.StatusOK, response.Success(ctx))
}
