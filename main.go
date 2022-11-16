package main

import (
	"GO_Rest_Api/controllers"
	"GO_Rest_Api/task"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=127.0.0.1 user=postgres password=admin dbname=todoGO port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}
	db.AutoMigrate(&task.Task{})

	taskRepository := task.TaskRepository(db)
	taskUseCase := task.TaskUseCase(taskRepository)
	taskControllers := controllers.TaskController(taskUseCase)

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/", taskControllers.TodoList)
	v1.GET("/:id", taskControllers.GetTaskById)
	v1.POST("/add", taskControllers.CreateTask)

	router.Run(":3000")
}
