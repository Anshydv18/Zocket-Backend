package main

import (
	"backend/base"
	"backend/middlewares"
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middlewares.SetContext())
	base.ConnectDB()
	routes.InitiateRoutes(router)

	router.Run(":8000")
}
