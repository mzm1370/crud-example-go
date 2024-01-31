package main

import (
	"crud-example-go/api"
	"crud-example-go/database"
	"crud-example-go/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Init()
	router := gin.Default()
	router.Use(middleware.LoggerMiddleware())
	api.SetupRoutes(router)
	router.Run(":8080")
}
