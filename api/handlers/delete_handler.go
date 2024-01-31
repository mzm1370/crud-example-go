package handlers

import (
	"crud-example-go/database"
	"crud-example-go/database/models"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
)

func DeleteUserHandler(c *gin.Context) {
	userID := c.Param("id")

	var user models.User

	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return

	}

	c.JSON(http.StatusOK, gin.H{"message": "USer deleted successfully"})
}
