package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/1_ToDoCRUD/config"
	"github.com/itsharshitk/1_ToDoCRUD/routes"
	"github.com/itsharshitk/1_ToDoCRUD/utils"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // swagger handler

	_ "github.com/itsharshitk/1_ToDoCRUD/docs" // generated docs
)

// @title ToDo CRUD API
// @version 1.0
// @description This is a sample server for a ToDo CRUD app.
// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey  BearerAuth
// @in                          header
// @name                        Authorization
// @description                 Type "Bearer <your-jwt-token>" to authenticate.
func main() {
	r := gin.Default()
	utils.InitValidations()
	config.ConnectDB()
	routes.ApiRoutes(r)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8000")
}
