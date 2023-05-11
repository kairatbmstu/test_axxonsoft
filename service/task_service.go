package service

import (
	"log"

	"example.com/test_axxonsoft/v2/database"
	"example.com/test_axxonsoft/v2/domain"
	"example.com/test_axxonsoft/v2/dto"
	"example.com/test_axxonsoft/v2/external"
	"example.com/test_axxonsoft/v2/repository"
	"github.com/gofrs/uuid"
)

type TaskService struct {
	TaskRepository   *repository.TaskRepository
	HeaderRepository *repository.HeaderRepository
	RabbitContext    *RabbitContext
	TaskMapper       TaskMapper
	TaskClient       external.TaskClient
}

func (t *TaskService) GetById(id uuid.UUID) (*dto.TaskDTO, error) {

	tx, err := database.DB.Begin()
	if err != nil {
		log.Println("error while opening transaction taskRepository.getById() method: ", err)
		return nil, err
	}
	task, err := t.TaskRepository.GetById(tx, id)

	if err != nil {
		log.Println("error calling taskRepository.getById() method: ", err)
		tx.Rollback()
		return nil, err
	}

	requestHeaders, err := t.HeaderRepository.GetRequestHeaders(tx, id)
	if err != nil {
		log.Println("error calling taskRepository.GetRequestHeaders() method: ", err)
		tx.Rollback()
		return nil, err
	}
	task.RequestHeaders = *requestHeaders

	responseHeaders, err := t.HeaderRepository.GetResponseHeaders(tx, id)
	if err != nil {
		log.Println("error calling taskRepository.GetResponseHeaders() method: ", err)
		tx.Rollback()
		return nil, err
	}
	task.ResponseHeaders = *responseHeaders

	err = tx.Commit()
	if err != nil {
		log.Println("error while commiting transaction in taskRepository.getById() method : ", err)
		return nil, err
	}

	var taskDto = t.TaskMapper.MapToDto(task)
	return taskDto, nil
}

// Create New Task entity
func (c *TaskService) CreateNewTask(taskDTO *dto.TaskDTO) (*dto.TaskDTO, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		log.Println("error while opening transaction taskRepository.getById() method: ", err)
		return nil, err
	}

	taskDTO.TaskStatus = domain.TaskStatusNew
	task := c.TaskMapper.MapToEntity(taskDTO)
	err = c.TaskRepository.Create(tx, task)

	if err != nil {
		log.Println("error calling taskRepository.getById() method: ", err)
		tx.Rollback()
		return nil, err
	}

	for _, header := range task.RequestHeaders {
		err := c.HeaderRepository.Create(tx, &header)
		if err != nil {
			log.Println("error calling taskRepository.getById() method: ", err)
			tx.Rollback()
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println("error while commiting transaction in taskRepository.CreateNewTask() method : ", err)
		return nil, err
	}

	var taskDto = c.TaskMapper.MapToDto(task)

	return taskDto, nil
}

func (c *TaskService) SendToQueue(taskDTO *dto.TaskDTO) (*dto.TaskDTO, error) {

	tx, err := database.DB.Begin()
	if err != nil {
		log.Println("error while opening transaction taskRepository.getById() method: ", err)
		return nil, err
	}

	// save and change task status to in_process
	taskDTO.TaskStatus = domain.TaskStatusInProcess
	var task = c.TaskMapper.MapToEntity(taskDTO)
	c.TaskRepository.Update(tx, task)

	err = tx.Commit()
	if err != nil {
		log.Println("error while commiting transaction in taskRepository.CreateNewTask() method : ", err)
		return nil, err
	}
	// send taskDto to Queue
	c.RabbitContext.SendTask(taskDTO)
	return taskDTO, nil
}

func (c *TaskService) ReceiveFromQueue(taskDTO *dto.TaskDTO) error {
	log.Println("received taskDto : ", taskDTO)
	//receive taskDto
	//make http request
	//handle response from http request
	//if http ~ 4xx or 5xx
	//save response with task status error
	//if http status ~ 2xx
	//save response with task status done
	var request = external.HttpRequest{
		Method:         taskDTO.Method,
		Url:            taskDTO.Url,
		RequestBody:    taskDTO.RequestBody,
		RequestHeaders: taskDTO.RequestHeaders,
	}
	response, err := c.TaskClient.DoHttpRequest(&request)
	if err != nil {
		log.Println(" Error occured when doing http request : ", err.Error())
		taskDTO.TaskStatus = domain.TaskStatusError
		return err
	}

	taskDTO.ResponseBody = response.ResponseBody
	taskDTO.ResponseHeaders = response.ResponseHeaders
	taskDTO.HttpStatusCode = response.Status
	taskDTO.TaskStatus = domain.TaskStatusDone
	c.Update(taskDTO)
	return nil
}

// Saves Response received from http request
func (c *TaskService) Update(taskDTO *dto.TaskDTO) (*dto.TaskDTO, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		log.Println("error while opening transaction taskRepository.getById() method: ", err)
		return nil, err
	}

	taskEntity := c.TaskMapper.MapToEntity(taskDTO)
	err = c.TaskRepository.Update(tx, taskEntity)

	if err != nil {
		log.Println("error calling taskRepository.getById() method: ", err)
		err := tx.Rollback()
		if err != nil {
			log.Println("error while rolling back transaction in taskRepository.getById() method : ", err)
		}
		return nil, err
	}

	for _, header := range taskEntity.ResponseHeaders {
		err := c.HeaderRepository.Create(tx, &header)
		if err != nil {
			log.Println("error calling taskRepository.getById() method: ", err)
			err := tx.Rollback()
			if err != nil {
				log.Println("error while rolling back transaction in taskRepository.getById() method : ", err)
			}
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println("error while commiting transaction in taskRepository.getById() method : ", err)
		return nil, err
	}

	var taskDto = c.TaskMapper.MapToDto(taskEntity)
	return taskDto, nil
}

func (c *TaskService) ChangeTaskStatus(taskDTO *dto.TaskDTO) error {

	tx, err := database.DB.Begin()
	if err != nil {
		log.Println("error while opening transaction taskRepository.getById() method: ", err)
		return err
	}

	taskEntity := c.TaskMapper.MapToEntity(taskDTO)
	err = c.TaskRepository.ChangeTaskStatus(tx, taskEntity.Id, taskEntity.TaskStatus)

	if err != nil {
		log.Println("error while commiting transaction in taskRepository.getById() method : ", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Println("error while commiting transaction in taskRepository.getById() method : ", err)
		return err
	}

	return nil
}

// TaskMappers = mappers entity to dto, and dto to entity object
type TaskMapper struct {
}

func (tm *TaskMapper) MapToDto(task *domain.Task) *dto.TaskDTO {
	var taskDto = dto.TaskDTO{
		Id:             task.Id,
		Method:         task.Method,
		Url:            task.Url,
		HttpStatusCode: task.HttpStatusCode,
		TaskStatus:     task.TaskStatus,
		ResponseLength: task.ResponseLength,
		RequestBody:    task.RequestBody,
		ResponseBody:   task.ResponseBody,
	}

	taskDto.RequestHeaders = make(map[string]string)
	for _, header := range task.RequestHeaders {
		taskDto.RequestHeaders[header.Name] = header.Value
	}

	taskDto.ResponseHeaders = make(map[string]string)
	for _, header := range task.ResponseHeaders {
		taskDto.ResponseHeaders[header.Name] = header.Value
	}

	return &taskDto
}

func (tm *TaskMapper) MapToEntity(taskDTO *dto.TaskDTO) *domain.Task {
	var task = domain.Task{
		HttpStatusCode: taskDTO.HttpStatusCode,
		Method:         taskDTO.Method,
		RequestBody:    taskDTO.RequestBody,
		ResponseBody:   taskDTO.ResponseBody,
		ResponseLength: taskDTO.ResponseLength,
		TaskStatus:     taskDTO.TaskStatus,
		Url:            taskDTO.Url,
	}

	for headerName, headerValue := range taskDTO.RequestHeaders {
		var header = domain.Header{
			Name:  headerName,
			Value: headerValue,
		}
		task.RequestHeaders = append(task.RequestHeaders, header)
	}

	for headerName, headerValue := range taskDTO.ResponseHeaders {
		var header = domain.Header{
			Name:  headerName,
			Value: headerValue,
		}
		task.ResponseHeaders = append(task.ResponseHeaders, header)
	}

	return &task
}
