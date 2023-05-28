package domain

import (
	"github.com/google/uuid"
)

/*
Package domain provides the domain models and types for managing tasks and headers.
*/

/*
The Task struct represents a task with its attributes such as the unique identifier (Id),
the HTTP method (Method), the associated URL (Url), the HTTP status code (HttpStatusCode)
returned by the task, the current status (TaskStatus) of the task, the response body length (ResponseLength),
the request headers (RequestHeaders), the request body (RequestBody),
the response headers (ResponseHeaders), and the response body (ResponseBody).
*/
type Task struct {
	Id              uuid.UUID  // The unique identifier of the task.
	Method          string     // The HTTP method used for the task.
	Url             string     // The URL associated with the task.
	HttpStatusCode  int        // The HTTP status code returned by the task.
	TaskStatus      TaskStatus // The current status of the task.
	ResponseLength  int        // The length of the response body.
	RequestHeaders  []Header   // The headers included in the request.
	RequestBody     string     // The body of the request.
	ResponseHeaders []Header   // The headers included in the response.
	ResponseBody    string     // The body of the response.
}

/**
The Header struct represents a header with its attributes such as the unique identifier (Id),
the ID of the task associated with the request header (RequestTaskId), the ID of the task associated
 with the response header (ResponseTaskId), the name of the header (Name), and the value of the header (Value).
*/
type Header struct {
	Id             uuid.UUID  // The unique identifier of the header.
	RequestTaskId  *uuid.UUID // The ID of the task associated with the request header.
	ResponseTaskId *uuid.UUID // The ID of the task associated with the response header.
	Name           string     // The name of the header.
	Value          string     // The value of the header.
}

/*
The TaskStatus type is a string type that represents the status of a task.
It can take one of the following constants: TaskStatusNew, TaskStatusInProcess, TaskStatusDone, or TaskStatusError.
*/
type TaskStatus string

const (
	TaskStatusNew       TaskStatus = "new"        // Indicates that the task is new and not processed yet.
	TaskStatusInProcess TaskStatus = "in_process" // Indicates that the task is currently being processed.
	TaskStatusDone      TaskStatus = "done"       // Indicates that the task has been successfully completed.
	TaskStatusError     TaskStatus = "error"      // Indicates that an error occurred while processing the task.
)
