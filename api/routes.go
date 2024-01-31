package api

import (
	"github.com/gin-gonic/gin"

	"crud-example-go/api/handlers"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/api/get_user/:id", handlers.GetUserHandler)
	router.POST("/api/create_user", handlers.CreateUserHandler)
	router.DELETE("/api/delete_user/:id", handlers.DeleteUserHandler)
	router.PUT("/api/update_user/:id", handlers.UpdateUserHandler)
}
