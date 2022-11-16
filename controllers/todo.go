package controllers

import (
	"fmt"
	"net/http"

	"GO_Rest_Api/task"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type taskController struct {
	taskUseCase task.UseCaseTask
}

func TaskController(taskUseCase task.UseCaseTask) *taskController {

	return &taskController{taskUseCase}
}

func (controller *taskController) TodoList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (controller *taskController) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (controller *taskController) CreateTask(c *gin.Context) {
	var taskRequest task.TaskRequest

	err := c.ShouldBindJSON(&taskRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf(e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	task, err := controller.taskUseCase.Create(taskRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erros": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})

}
