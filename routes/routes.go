package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/1_ToDoCRUD/controllers"
	"github.com/itsharshitk/1_ToDoCRUD/middleware"
)

func ApiRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	v1 := r.Group("/", middleware.AuthMiddleware())
	{
		v1.GET("/profile", controllers.UserProfile)
		v1.POST("/task", controllers.AddTask)
		v1.GET("/task", controllers.GetTasks)
		v1.PUT("/task", controllers.UpdateTask)
		v1.DELETE("/task/:id", controllers.DeleteTask)
	}
}
