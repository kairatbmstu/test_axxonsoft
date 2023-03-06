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
	TaskStatus_NEW        TaskStatus = "new"
	TaskStatus_IN_PROCESS TaskStatus = "in_process"
	TaskStatus_DONE       TaskStatus = "done"
	TaskStatus_ERROR      TaskStatus = "error"
)
