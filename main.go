package main

import (
	"rest/goapi/endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/users", endpoints.GetAllUsers)
	router.POST("/users", endpoints.AddUser)
	router.GET("/users/:id", endpoints.GetUserByID)
	router.Run("localhost:8080")
}
