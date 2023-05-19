package dto

import (
	"example.com/test_axxonsoft/v2/domain"
	"github.com/gofrs/uuid"
)

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

type HeaderDTO struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TaskIdDTO struct {
	Id uuid.UUID `json:"id"`
}

type ErrorDTO struct {
	HasErrors bool
	Errors    []string
}
