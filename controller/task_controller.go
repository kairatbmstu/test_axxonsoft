package controller

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"example.com/test_axxonsoft/v2/domain"
	"example.com/test_axxonsoft/v2/dto"
	"example.com/test_axxonsoft/v2/service"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type TaskController struct {
	TaskService   *service.TaskService
	TaskValidator TaskValidator
}

func (t *TaskController) PostTask(c *gin.Context) {
	var taskDto = dto.TaskDTO{}
	if err := c.ShouldBindJSON(&taskDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":        "bad_request",
			"errorMessage": err.Error(),
		})
		return
	}

	var error = t.TaskValidator.validate(&taskDto)

	if error.HasErrors {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": error.Errors,
		})
		return
	}

	taskResultDto, err := t.TaskService.Create(&taskDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":        "internal_server_error",
			"errorMessage": err.Error(),
		})
		return
	}

	taskResultDto, err = t.TaskService.SendToQueue(taskResultDto)

	if err != nil {
		log.Println("error occured when sendToQueue : ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":        "internal_server_error",
			"errorMessage": err.Error(),
		})
		return
	}

	taskResultDto.TaskStatus = domain.TaskStatusInProcess

	err = t.TaskService.ChangeTaskStatus(taskResultDto)

	if err != nil {
		log.Println("error occured when change taskDto : ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":        "internal_server_error",
			"errorMessage": err.Error(),
		})
		return
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

type TaskValidator struct {
}

func (t TaskValidator) validate(taskDto *dto.TaskDTO) *dto.ErrorDTO {
	taskDto.Method = strings.TrimSpace(taskDto.Method)

	var errors = new(dto.ErrorDTO)
	errors.HasErrors = false
	if len(taskDto.Method) == 0 {
		errors.Errors = append(errors.Errors, "Method field is required")
		errors.HasErrors = true
	}

	if len(taskDto.Method) > 0 {
		if !(strings.EqualFold(taskDto.Method, "GET") || strings.EqualFold(taskDto.Method, "POST") ||
			strings.EqualFold(taskDto.Method, "PUT") || strings.EqualFold(taskDto.Method, "DELETE") ||
			strings.EqualFold(taskDto.Method, "HEAD")) {
			errors.Errors = append(errors.Errors, "method is not allowed")
			errors.HasErrors = true
		}
	}

	if !IsUrl(taskDto.Url) {
		errors.Errors = append(errors.Errors, "URL is not valid")
		errors.HasErrors = true
	}

	if errors.HasErrors {
		return errors
	}
	return errors
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
