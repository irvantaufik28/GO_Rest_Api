package main

import (
	"GO_Rest_Api/controllers"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=127.0.0.1 user=postgres password=admin dbname=todoGO port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}
	fmt.Println("Database connction......")

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/", controllers.TodoList)
	v1.GET("/:id", controllers.GetTaskById)
	v1.POST("/add", controllers.CreateTask)

	router.Run(":3000")
}
