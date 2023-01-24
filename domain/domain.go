package domain

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
	RequestTaskId   *string
	ResponsetTaskId *string
	Name            string
	Value           string
}

type TaskStatus string

const (
	NEW        TaskStatus = "new"
	IN_PROCESS TaskStatus = "in_process"
	DONE       TaskStatus = "done"
	ERROR      TaskStatus = "error"
)
