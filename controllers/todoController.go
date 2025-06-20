package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/1_ToDoCRUD/config"
	"github.com/itsharshitk/1_ToDoCRUD/model"
	"gorm.io/gorm"
)

func AddTask(c *gin.Context) {
	var tasks model.Todo
	if err := c.ShouldBindJSON(&tasks); err != nil {
		c.JSON(http.StatusBadRequest, &model.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Input: " + err.Error(),
		})
		return
	}
	tasks.UserId = c.GetUint("id")

	if err := config.Db.Create(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &model.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Something Went Wrong: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Id":      tasks.ID,
		"Message": "Task Added Successfully",
	})
}

func GetTasks(c *gin.Context) {
	var tasks []model.Todo
	id := c.GetUint("id")

	result := config.Db.Where("user_id = ?", id).Find(&tasks).Error

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, &model.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong: " + result.Error(),
		})
		return
	}

	// If no tasks are found, you can return an empty array or a specific message
	if len(tasks) == 0 {
		c.JSON(http.StatusOK, []model.Todo{}) // Return an empty array
		return
	}

	c.JSON(http.StatusOK, tasks)

}

func TaskById(c *gin.Context) {
	taskId := c.Param("id")
	userId := c.GetUint("id")
	var task model.Todo

	result := config.Db.Where("id = ?", taskId).Where("user_id = ?", userId).First(&task)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, &model.ErrorResponse{
				Status:  http.StatusNotFound,
				Message: "Task not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, &model.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong: " + result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)

}

func UpdateTask(c *gin.Context) {
	var task model.Todo

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, &model.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Input: " + err.Error(),
		})
		return
	}

	if task.IsComplete == nil {
		result := config.Db.Model(&task).Where("id = ?", task.ID).Where("user_id = ?", c.GetUint("id")).Updates(map[string]any{"title": task.Title, "description": task.Description})
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, &model.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to update task: " + result.Error.Error(),
			})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, &model.ErrorResponse{
				Status:  http.StatusNotFound,
				Message: "Task not found or no changes made",
			})
			return
		}
	} else {
		result := config.Db.Model(&task).Where("id = ?", task.ID).Where("user_id = ?", c.GetUint("id")).Updates(map[string]any{"title": task.Title, "description": task.Description, "is_complete": task.IsComplete})
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, &model.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to update task: " + result.Error.Error(),
			})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, &model.ErrorResponse{
				Status:  http.StatusNotFound,
				Message: "Task not found or no changes made",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Task updated successfully",
	})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task model.Todo

	result := config.Db.Delete(&task, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, &model.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to Delete Task: " + result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, &model.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Task not found or already deleted",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Task Deleted Successfully",
	})
}
