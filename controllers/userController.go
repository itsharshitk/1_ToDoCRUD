package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/1_ToDoCRUD/config"
	"github.com/itsharshitk/1_ToDoCRUD/model"
)

func UserProfile(c *gin.Context) {
	userId := c.GetUint("id")

	var user model.User

	result := config.Db.Where("id = ?", userId).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "User Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, model.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}
