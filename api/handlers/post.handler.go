package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"crud-example-go/database"
	"crud-example-go/database/models"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Tel      string `json:"tel"`
}

func CreateUserHandler(c *gin.Context) {
	var request CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := models.User{
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
		Address:  request.Address,
		Tel:      request.Tel,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, newUser)

}
