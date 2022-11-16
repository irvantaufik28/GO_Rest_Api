package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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
	tasks, err := controller.taskUseCase.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	var tasksResponse []task.TaskResponse
	for _, t := range tasks {
		taskResponse := task.TaskResponse{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Doing:       t.Doing,
		}
		tasksResponse = append(tasksResponse, taskResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": tasksResponse,
	})

}

func (controller *taskController) GetTaskById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	t, err := controller.taskUseCase.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	taskResponse := task.TaskResponse{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Doing:       t.Doing,
	}
	c.JSON(http.StatusOK, gin.H{
		"data": taskResponse,
	})
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
