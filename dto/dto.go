package dto

import "example.com/test_axxonsoft/v2/domain"

type TaskDTO struct {
	Id              string            `json:"id"`
	Method          string            `json:"method"`
	Url             string            `json:"url"`
	HttpStatusCode  string            `json:"httpStatusCode"`
	TaskStatus      domain.TaskStatus `json:"taskStatus"`
	ResponseLength  int               `json:"responseLength"`
	RequestHeaders  []HeaderDTO       `json:"requestHeaders"`
	RequestBody     string            `json:"requestBody"`
	ResponseHeaders []HeaderDTO       `json:"responseHeaders"`
	ResponseBody    string            `json:"responseBody"`
}

type HeaderDTO struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
