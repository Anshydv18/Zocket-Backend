package routes

import (
	"backend/handler"

	"github.com/gin-gonic/gin"
)

func InitiateRoutes(router *gin.Engine) {
	ApiGroup := router.Group("/api")

	UserApi := ApiGroup.Group("/user")
	UserApi.POST("/login", handler.UserLogin)
	UserApi.POST("/createUser", handler.CreateUserProfile)
}
