package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/1_ToDoCRUD/config"
	"github.com/itsharshitk/1_ToDoCRUD/model"
	"gorm.io/gorm"
)

// AddTaskTest godoc
// @Summary      Add new Task
// @Description  Adds a new task by title and description
// @Tags         ToDo
// @Accept       json
// @Produce      json
// @Param        tasks  body      model.AddTaskRequest  true  "Todo data"
// @Success      201   {object}  model.APIResponse
// @Failure      500   {object}  model.APIResponse
// @Security     BearerAuth
// @Router       /task [post]
func AddTask(c *gin.Context) {
	var tasks model.AddTaskRequest
	if err := c.ShouldBindJSON(&tasks); err != nil {
		c.JSON(http.StatusBadRequest, &model.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Input",
			Details: err.Error(),
		})
		return
	}

	newTask := model.Todo{
		UserId:      c.GetUint("id"),
		Title:       tasks.Title,
		Description: tasks.Description,
	}

	if err := config.Db.Create(&newTask).Error; err != nil {
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
		Data:    newTask,
	})
}

// GetTasks godoc
// @Summary      Get user's tasks
// @Description  Fetch all tasks of a user
// @Tags         ToDo
// @Produce      json
// @Success      200   {object}  model.Todo
// @Failure      500   {object}  model.APIResponse
// @Security     BearerAuth
// @Router       /task [get]
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

	// If no tasks are found, you can return an empty array or a specific message
	if len(tasks) == 0 {
		c.JSON(http.StatusOK, []model.Todo{}) // Return an empty array
		return
	}

	c.JSON(http.StatusOK, tasks)

}

// TasksById godoc
// @Summary      Get Task by Id
// @Description  Returns a single task for the given ID belonging to the user
// @Tags         ToDo
// @Produce      json
// @Param        Id path int true "Task ID"
// @Success      200   {object}  model.Todo
// @Failure      400   {object}  model.APIResponse
// @Failure      500   {object}  model.APIResponse
// @Security     BearerAuth
// @Router       /task/{id} [get]
func TasksById(c *gin.Context) {
	userId := c.GetUint("id")
	taskId := c.Param("id")
	var task model.Todo

	result := config.Db.Where("id = ? AND user_id = ?", taskId, userId).First(&task)

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

// UpdateTask godoc
// @Summary      Update existing task
// @Description  Update a task's title, description, and completion status
// @Tags         ToDo
// @Accept       json
// @Produce      json
// @Param        id    path      int        true  "Task ID"
// @Param        task  body      model.Todo  true  "Updated task data"
// @Success      200   {object}  model.APIResponse
// @Failure      400   {object}  model.APIResponse
// @Failure      404   {object}  model.APIResponse
// @Failure      500   {object}  model.APIResponse
// @Security     BearerAuth
// @Router       /task/{id} [put]
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

// DeleteTask godoc
// @Summary      Delete task
// @Description  Deletes a task by its ID
// @Tags         ToDo
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Success      200  {object}  model.APIResponse
// @Failure      404  {object}  model.APIResponse
// @Failure      500  {object}  model.APIResponse
// @Security     BearerAuth
// @Router       /task/{id} [delete]
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
