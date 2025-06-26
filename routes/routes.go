package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/1_ToDoCRUD/controllers"
	"github.com/itsharshitk/1_ToDoCRUD/middleware"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // swagger handler

	_ "github.com/itsharshitk/1_ToDoCRUD/docs" // generated docs
)

func ApiRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	// Swagger route
	r.GET("/swagger", func(c *gin.Context) {
		c.Redirect(301, "/swagger/index.html")
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/", middleware.AuthMiddleware())
	{
		v1.GET("/profile", controllers.UserProfile)
		v1.POST("/task", controllers.AddTask)
		v1.GET("/task", controllers.GetTasks)
		v1.GET("/task/:id", controllers.TasksById)
		v1.PUT("/task/:id", controllers.UpdateTask)
		v1.DELETE("/task/:id", controllers.DeleteTask)
	}
}
