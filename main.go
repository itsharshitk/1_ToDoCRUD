package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/1_ToDoCRUD/config"
	"github.com/itsharshitk/1_ToDoCRUD/routes"
	"github.com/itsharshitk/1_ToDoCRUD/utils"
)

// @title ToDo CRUD API
// @version 1.1
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

	r.Run(":8000")
}
