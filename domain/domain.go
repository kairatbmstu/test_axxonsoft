package domain

import (
	"github.com/gofrs/uuid"
)

type Task struct {
	Id              uuid.UUID
	Method          string
	Url             string
	HttpStatusCode  int
	TaskStatus      TaskStatus
	ResponseLength  int
	RequestHeaders  []Header
	RequestBody     string
	ResponseHeaders []Header
	ResponseBody    string
}

type Header struct {
	Id             uuid.UUID
	RequestTaskId  *uuid.UUID
	ResponseTaskId *uuid.UUID
	Name           string
	Value          string
}

type TaskStatus string

const (
	TaskStatusNew       TaskStatus = "new"
	TaskStatusInProcess TaskStatus = "in_process"
	TaskStatusDone      TaskStatus = "done"
	TaskStatusError     TaskStatus = "error"
)
