package dto

import "example.com/test_axxonsoft/v2/domain"

type TaskDTO struct {
	Id              string
	Method          string
	Url             string
	HttpStatusCode  string
	TaskStatus      domain.TaskStatus
	ResponseLength  int
	RequestHeaders  []HeaderDTO
	RequestBody     string
	ResponseHeaders []HeaderDTO
	ResponseBody    string
}

type HeaderDTO struct {
	Name  string
	Value string
}
