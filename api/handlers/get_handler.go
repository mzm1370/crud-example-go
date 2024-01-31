package handlers

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"github.com/jinzhu/gorm"

	"crud-example-go/database"
	"crud-example-go/database/models"
)

func GetUserHandler(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {

		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internet Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}
