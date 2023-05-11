package controller

import (
	"fmt"
	"log"
	"net/http"

	"example.com/test_axxonsoft/v2/domain"
	"example.com/test_axxonsoft/v2/dto"
	"example.com/test_axxonsoft/v2/service"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type TaskController struct {
	TaskService   *service.TaskService
	RabbitContext *service.RabbitContext
}

func (t *TaskController) PostTask(c *gin.Context) {
	var taskDto = dto.TaskDTO{}
	if err := c.BindJSON(&taskDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":        "bad_request",
			"errorMessage": err.Error(),
		})
		return
	}

	taskResultDto, err := t.TaskService.CreateNewTask(&taskDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":        "internal_server_error",
			"errorMessage": err.Error(),
		})
	}

	taskResultDto, err = t.TaskService.SendToQueue(taskResultDto)

	if err != nil {
		log.Println("error occured when sendToQueue : ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":        "internal_server_error",
			"errorMessage": err.Error(),
		})
	}

	taskResultDto.TaskStatus = domain.TaskStatusInProcess

	err = t.TaskService.ChangeTaskStatus(taskResultDto)

	if err != nil {
		log.Println("error occured when change taskDto : ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":        "internal_server_error",
			"errorMessage": err.Error(),
		})
	}

	c.JSON(200, taskResultDto)
}

func (t *TaskController) GetTask(c *gin.Context) {
	var id = c.Params.ByName("id")
	fmt.Println("id : " + id)
	uid, err := uuid.FromString(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":        "bad_request",
			"errorMessage": err.Error(),
		})
	}
	task, err := t.TaskService.GetById(uid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":        "internal_server_error",
			"errorMessage": err.Error(),
		})
	}

	c.JSON(200, task)
}
