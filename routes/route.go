package routes

import (
	"backend/handler"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func InitiateRoutes(router *gin.Engine) {
	ApiGroup := router.Group("/api")

	UserApi := ApiGroup.Group("/user")
	UserApi.POST("/logout", handler.LogoutUser)
	UserApi.POST("/login", handler.UserLogin)
	UserApi.POST("/createUser", handler.CreateUserProfile)

	ProtectedApi := ApiGroup.Group("/p1")
	ProtectedApi.Use(middlewares.AuthenticateUser())
	ProtectedApi.POST("/createTask", handler.CreateTask)
	ProtectedApi.GET("fetchAllTask", handler.FetchTasks)
	ProtectedApi.PATCH("/updateTask", handler.UpdateTask)
	ProtectedApi.POST("/fetchTaskName", handler.GenerateTaskName)

}
