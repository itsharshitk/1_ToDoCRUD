package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/1_ToDoCRUD/config"
	"github.com/itsharshitk/1_ToDoCRUD/routes"
)

func main() {
	r := gin.Default()
	config.ConnectDB()
	routes.ApiRoutes(r)
	r.Run(":8000")
}
