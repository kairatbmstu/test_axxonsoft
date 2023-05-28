package dto

import (
	"example.com/test_axxonsoft/v2/domain"
	"github.com/google/uuid"
)

/*
Package dto provides the data transfer objects (DTOs) for transferring data between different layers or components of an application.
*/

/*
The TaskDTO struct represents a data transfer object for a task. It includes attributes such
as the task ID (Id), the HTTP method (Method), the URL (Url), the HTTP status code (HttpStatusCode),
the task status (TaskStatus), the response body length (ResponseLength), the request headers (RequestHeaders),
the request body (RequestBody), the response headers (ResponseHeaders), and the response body (ResponseBody).
This struct is designed to facilitate data transfer in JSON format.
*/
type TaskDTO struct {
	Id              uuid.UUID         `json:"id"`
	Method          string            `json:"method" binding:"required"`
	Url             string            `json:"url" binding:"required"`
	HttpStatusCode  int               `json:"httpStatusCode"`
	TaskStatus      domain.TaskStatus `json:"taskStatus"`
	ResponseLength  int               `json:"responseLength"`
	RequestHeaders  map[string]string `json:"requestHeaders"`
	RequestBody     string            `json:"requestBody"`
	ResponseHeaders map[string]string `json:"responseHeaders"`
	ResponseBody    string            `json:"responseBody"`
}

/*
The HeaderDTO struct represents a data transfer object for a header.
It includes attributes such as the header name (Name) and the header value (Value).
This struct is used to transfer header data between different components of the application.
*/
type HeaderDTO struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

/*
The TaskIdDTO struct represents a data transfer object for a task ID.
It includes a single attribute, the task ID (Id), which is used for transferring task ID information.
*/
type TaskIdDTO struct {
	Id uuid.UUID `json:"id"`
}

/*
The ErrorDTO struct represents a data transfer object for handling errors.
It includes attributes such as a boolean flag indicating whether errors exist (HasErrors) and a slice of error messages (Errors).
*/
type ErrorDTO struct {
	HasErrors bool
	Errors    []string
}

/*
The TaskStatusDTO struct represents a data transfer object for task status.
It includes attributes similar to the TaskDTO struct, such as the task ID (Id),
the HTTP method (Method), the URL (Url), the HTTP status code (HttpStatusCode),
the task status (TaskStatus), the response body length (ResponseLength),
the request headers (RequestHeaders), the request body (RequestBody), and
the response headers (ResponseHeaders). This struct is designed to transfer task status information.
*/
type TaskStatusDTO struct {
	Id              uuid.UUID         `json:"id"`
	Method          string            `json:"method" binding:"required"`
	Url             string            `json:"url" binding:"required"`
	HttpStatusCode  int               `json:"httpStatusCode"`
	TaskStatus      domain.TaskStatus `json:"taskStatus"`
	ResponseLength  int               `json:"responseLength"`
	RequestHeaders  map[string]string `json:"requestHeaders"`
	RequestBody     string            `json:"requestBody"`
	ResponseHeaders map[string]string `json:"responseHeaders"`
}
