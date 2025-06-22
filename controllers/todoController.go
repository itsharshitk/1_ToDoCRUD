package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/1_ToDoCRUD/config"
	"github.com/itsharshitk/1_ToDoCRUD/model"
	"gorm.io/gorm"
)

func AddTask(c *gin.Context) {
	var tasks model.Todo
	if err := c.ShouldBindJSON(&tasks); err != nil {
		c.JSON(http.StatusBadRequest, &model.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Input",
			Details: err.Error(),
		})
		return
	}
	tasks.UserId = c.GetUint("id")

	if err := config.Db.Create(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &model.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Something Went Wrong",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &model.APIResponse{
		Status:  http.StatusCreated,
		Message: "Task Added Successfully",
		Data:    map[string]uint{"task_id": tasks.ID},
	})
}

func GetTasks(c *gin.Context) {
	var tasks []model.Todo
	id := c.GetUint("id")

	if err := config.Db.Where("user_id = ?", id).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &model.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Something Went Wrong",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, tasks)

}

func TasksById(c *gin.Context) {
	userId := c.GetUint("id")
	taskId := c.Param("id")
	var task model.Todo

	result := config.Db.Where("id = ?", taskId).Where("user_id = ?", userId).First(&task)

	if result.Error != nil {
		if gorm.ErrRecordNotFound != nil {
			c.JSON(http.StatusBadRequest, &model.APIResponse{
				Status:  http.StatusBadRequest,
				Message: "Task not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, &model.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong",
			Details: result.Error.Error(),
		})
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	var task model.Todo
	task_id := c.Param("id")

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, &model.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Input",
			Details: err.Error(),
		})
		return
	}

	if task.IsComplete == nil {
		result := config.Db.Model(&task).Where("id = ? AND user_id = ?", task_id, c.GetUint("id")).Updates(map[string]any{"title": task.Title, "description": task.Description})
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, &model.APIResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to update task",
				Details: result.Error.Error(),
			})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, &model.APIResponse{
				Status:  http.StatusNotFound,
				Message: "Task not found or no changes made",
			})
			return
		}
	} else {
		result := config.Db.Model(&task).Where("id = ? AND user_id = ?", task_id, c.GetUint("id")).Updates(map[string]any{"title": task.Title, "description": task.Description, "is_complete": task.IsComplete})
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, &model.APIResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to update task",
				Details: result.Error.Error(),
			})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, &model.APIResponse{
				Status:  http.StatusNotFound,
				Message: "Task not found or no changes made",
			})
			return
		}
	}

	c.JSON(http.StatusOK, &model.APIResponse{
		Status:  http.StatusOK,
		Message: "Task updated successfully",
	})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task model.Todo

	result := config.Db.Delete(&task, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, &model.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to Delete Task",
			Details: result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, &model.APIResponse{
			Status:  http.StatusNotFound,
			Message: "Task not found or already deleted",
		})
		return
	}

	c.JSON(http.StatusOK, &model.APIResponse{
		Status:  http.StatusOK,
		Message: "Task Deleted Successfully",
	})
}
