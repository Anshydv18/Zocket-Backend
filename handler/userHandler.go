package handler

import (
	"backend/models/requests"
	"backend/models/response"
	"backend/services"
	"backend/utils"
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
		c.JSON(http.StatusBadRequest, response.Fail(ctx, request))
		return
	}

	auth_token := utils.SetAuthToken(c, data.Name)
	c.SetCookie("auth_token", auth_token, 1800, "/", "", false, true)

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

func LogoutUser(c *gin.Context) {
	c.SetCookie(
		"auth_token", "", 3600, "/", "", false, true,
	)
	c.JSON(http.StatusOK, "logout successful")
}
