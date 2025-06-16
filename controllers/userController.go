package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/1_ToDoCRUD/config"
	"github.com/itsharshitk/1_ToDoCRUD/model"
)

func UserProfile(c *gin.Context) {
	userId := c.GetUint("userId")

	var user model.User

	result := config.Db.Where("id = ?", userId).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User Not Found"})
	}

	updatedDate := user.UpdatedAt.Format("01-02-2006")

	c.JSON(http.StatusOK, gin.H{
		"id":           user.ID,
		"name":         user.Name,
		"email":        user.Email,
		"last_updated": updatedDate,
	})
}
