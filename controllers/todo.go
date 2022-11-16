package controllers

import (
	"fmt"
	"net/http"

	"GO_Rest_Api/task"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func TodoList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Irvan Taufik",
		"Job":  " Programmer",
	})
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func CreateTask(c *gin.Context) {
	var bookInput task.BookInput

	err := c.ShouldBindJSON(&bookInput)
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

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})

}
