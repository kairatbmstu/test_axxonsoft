package controller

import (
	"fmt"
	"net/http"

	"example.com/test_axxonsoft/v2/dto"
	"example.com/test_axxonsoft/v2/service"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func PostTask(c *gin.Context) {
	var taskDto = dto.TaskDTO{}
	if err := c.BindJSON(&taskDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":        "bad_request",
			"errorMessage": err.Error(),
		})
		return
	}

	taskResultDto, err := service.TaskServiceInst.CreateNewTask(taskDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":        "internal_server_error",
			"errorMessage": err.Error(),
		})
	}

	c.JSON(200, taskResultDto)
}

func GetTask(c *gin.Context) {
	var id = c.GetString("id")
	fmt.Println("id : " + id)
	uid, err := uuid.FromString(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":        "bad_request",
			"errorMessage": err.Error(),
		})
	}
	task, err := service.TaskServiceInst.GetById(uid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":        "internal_server_error",
			"errorMessage": err.Error(),
		})
	}

	c.JSON(200, task)
}
