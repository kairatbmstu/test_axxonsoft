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
	"github.com/google/uuid"
)

/*
Package controller provides the implementation of the controller layer responsible
for handling incoming HTTP requests and interacting with the domain and service layers.
*/
type TaskController struct {
	TaskService   *service.TaskService
	TaskValidator TaskValidator
}

/*
The PostTask method is an HTTP handler function that handles the HTTP POST request
for creating a new task. It retrieves the task data from the request body,
validates it using the TaskValidator, creates the task using the TaskService,
sends the task to a queue using the TaskService, changes the task status to
 "in_process" using the TaskService, and finally returns the task ID in the response body.
*/
func (t *TaskController) PostTask(c *gin.Context) {
	var taskDto = new(dto.TaskDTO)
	if err := c.ShouldBindJSON(&taskDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":        "bad_request",
			"errorMessage": err.Error(),
		})
		return
	}

	var error = t.TaskValidator.validate(taskDto)

	if error.HasErrors {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": error.Errors,
		})
		return
	}

	taskDto, err := t.TaskService.Create(taskDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":        "internal_server_error",
			"errorMessage": err.Error(),
		})
		return
	}

	taskDto, err = t.TaskService.SendToQueue(taskDto)

	if err != nil {
		log.Println("error occured when sendToQueue : ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":        "internal_server_error",
			"errorMessage": err.Error(),
		})
		return
	}

	taskDto.TaskStatus = domain.TaskStatusInProcess

	log.Println("taskDto.Id : {} ", taskDto.Id)
	err = t.TaskService.ChangeTaskStatus(taskDto)

	if err != nil {
		log.Println("error occured when change taskDto : ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":        "internal_server_error",
			"errorMessage": err.Error(),
		})
		return
	}

	var taskDtoId = dto.TaskIdDTO{
		Id: taskDto.Id,
	}

	c.JSON(200, taskDtoId)
}

/*
The GetTask method is an HTTP handler function that handles the HTTP GET request
for retrieving the status of a specific task. It extracts the task ID from the request URL,
retrieves the task status using the TaskService, creates a TaskStatusDTO object containing
the relevant task status information, and returns it in the response body.
*/
func (t *TaskController) GetTask(c *gin.Context) {
	var id = c.Params.ByName("id")
	fmt.Println("id : " + id)
	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":        "bad_request",
			"errorMessage": err.Error(),
		})
	}
	task, err := t.TaskService.GetTaskStatusById(uid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":        "internal_server_error",
			"errorMessage": err.Error(),
		})
	}

	var taskStatusDto = dto.TaskStatusDTO{
		Id:              task.Id,
		Method:          task.Method,
		Url:             task.Url,
		HttpStatusCode:  task.HttpStatusCode,
		RequestHeaders:  task.RequestHeaders,
		TaskStatus:      task.TaskStatus,
		ResponseHeaders: task.ResponseHeaders,
	}

	c.JSON(200, taskStatusDto)
}

/*
The validate method is a helper function of the TaskValidator struct.
It performs validation on the provided TaskDTO object and returns an
ErrorDTO object containing any validation errors encountered.
It checks for the presence of the Method field, validates the
allowed HTTP methods, and checks the validity of the URL.
*/
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

/*
The IsUrl function is a helper function that checks whether a given string represents a valid URL.
It uses the url.Parse function to parse the string and determines if the parsing was successful
based on the presence of a scheme and a host.
*/
func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
