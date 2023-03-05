package domain

import "database/sql"

type Task struct {
	Id              string
	Method          string
	Url             string
	HttpStatusCode  string
	TaskStatus      TaskStatus
	ResponseLength  int
	RequestHeaders  []Header
	RequestBody     string
	ResponseHeaders []Header
	ResponseBody    string
}

type Header struct {
	Id              string
	RequestTaskId   sql.NullString
	ResponsetTaskId sql.NullString
	Name            string
	Value           string
}

type TaskStatus string

const (
	TaskStatus_NEW        TaskStatus = "new"
	TaskStatus_IN_PROCESS TaskStatus = "in_process"
	TaskStatus_DONE       TaskStatus = "done"
	TaskStatus_ERROR      TaskStatus = "error"
)
