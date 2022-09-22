package main

import (
	"my-joke-book/database"
	"my-joke-book/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	database.ConnectDB()
	routes.UserRoute(router)
	router.Run(":8080")
}
