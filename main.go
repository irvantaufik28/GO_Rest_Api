package main

import (
	"GO_Rest_Api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/", controllers.TodoList)
	v1.GET("/:id", controllers.GetTaskById)
	v1.POST("/add", controllers.CreateTask)

	router.Run(":3000")
}
