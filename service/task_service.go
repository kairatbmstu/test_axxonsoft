package service

import (
	"log"
	"strings"

	"example.com/test_axxonsoft/v2/database"
	"example.com/test_axxonsoft/v2/domain"
	"example.com/test_axxonsoft/v2/dto"
	"example.com/test_axxonsoft/v2/repository"
	"github.com/gofrs/uuid"
)

var TaskServiceInst = TaskService{}

type TaskService struct {
	TaskRepository   repository.TaskRepository
	HeaderRepository repository.HeaderRepository
	TaskMapper       TaskMapper
}

func (t TaskService) GetById(id uuid.UUID) (*dto.TaskDTO, error) {

	tx, err := database.DB.Begin()
	if err != nil {
		log.Println("error while opening transaction taskRepository.getById() method: ", err)
		return nil, err
	}
	task, err := t.TaskRepository.GetById(tx, id)

	if err != nil {
		log.Println("error calling taskRepository.getById() method: ", err)
		err := tx.Rollback()
		if err != nil {
			log.Println("error while rolling back transaction in taskRepository.getById() method : ", err)
		}
		return nil, err
	}

	requestHeaders, err := t.HeaderRepository.GetRequestHeaders(tx, id)
	if err != nil {
		log.Println("error calling taskRepository.GetRequestHeaders() method: ", err)
		err := tx.Rollback()
		if err != nil {
			log.Println("error while rolling back transaction in taskRepository.getById() method : ", err)
		}
		return nil, err
	}
	task.RequestHeaders = *requestHeaders

	responseHeaders, err := t.HeaderRepository.GetResponseHeaders(tx, id)
	if err != nil {
		log.Println("error calling taskRepository.GetResponseHeaders() method: ", err)
		err := tx.Rollback()
		if err != nil {
			log.Println("error while rolling back transaction in taskRepository.getById() method : ", err)
		}
		return nil, err
	}
	task.ResponseHeaders = *responseHeaders

	err = tx.Commit()
	if err != nil {
		log.Println("error while commiting transaction in taskRepository.getById() method : ", err)
		return nil, err
	}

	var taskDto = t.TaskMapper.MapToDto(*task)
	return &taskDto, nil
}

func (c TaskService) Create(taskDTO dto.TaskDTO) (*dto.TaskDTO, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		log.Println("error while opening transaction taskRepository.getById() method: ", err)
		return nil, err
	}

	task := c.TaskMapper.MapToEntity(taskDTO)

	resultTask, err := c.TaskRepository.Create(tx, &task)

	if err != nil {
		log.Println("error calling taskRepository.getById() method: ", err)
		err := tx.Rollback()
		if err != nil {
			log.Println("error while rolling back transaction in taskRepository.getById() method : ", err)
		}
		return nil, err
	}

	var requestHeadersCreated = make([]domain.Header, 0)
	for _, header := range task.RequestHeaders {
		headerCreated, err := c.HeaderRepository.Create(tx, &header)
		if err != nil {
			log.Println("error calling taskRepository.getById() method: ", err)
			err := tx.Rollback()
			if err != nil {
				log.Println("error while rolling back transaction in taskRepository.getById() method : ", err)
			}
			return nil, err
		}

		requestHeadersCreated = append(requestHeadersCreated, *headerCreated)
	}

	resultTask.RequestHeaders = requestHeadersCreated

	var responseHeadersCreated = make([]domain.Header, 0)
	for _, header := range task.ResponseHeaders {
		headerCreated, err := c.HeaderRepository.Create(tx, &header)
		if err != nil {
			log.Println("error calling taskRepository.getById() method: ", err)
			err := tx.Rollback()
			if err != nil {
				log.Println("error while rolling back transaction in taskRepository.getById() method : ", err)
			}
			return nil, err
		}

		responseHeadersCreated = append(responseHeadersCreated, *headerCreated)
	}

	resultTask.ResponseHeaders = responseHeadersCreated

	err = tx.Commit()
	if err != nil {
		log.Println("error while commiting transaction in taskRepository.getById() method : ", err)
		return nil, err
	}

	var taskDto = c.TaskMapper.MapToDto(*resultTask)
	return &taskDto, nil
}

func (c TaskService) Update(taskDTO dto.TaskDTO) (*dto.TaskDTO, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		log.Println("error while opening transaction taskRepository.getById() method: ", err)
		return nil, err
	}

	taskEntity := c.TaskMapper.MapToEntity(taskDTO)
	resultTask, err := c.TaskRepository.Update(tx, &taskEntity)

	if err != nil {
		log.Println("error calling taskRepository.getById() method: ", err)
		err := tx.Rollback()
		if err != nil {
			log.Println("error while rolling back transaction in taskRepository.getById() method : ", err)
		}
		return nil, err
	}

	requestHeaders, err := c.HeaderRepository.GetRequestHeaders(tx, taskDTO.Id)

	if err != nil {
		log.Println("error calling taskRepository.getById() method: ", err)
		err := tx.Rollback()
		if err != nil {
			log.Println("error while rolling back transaction in taskRepository.getById() method : ", err)
		}
		return nil, err
	}

	for _, headerFromInput := range taskEntity.RequestHeaders {
		for _, headerFromDb := range *requestHeaders {
			if strings.EqualFold(headerFromInput.Name, headerFromDb.Name) {

			}
		}
	}

	responseHeaders, err := c.HeaderRepository.GetResponseHeaders(tx, taskDTO.Id)

	if err != nil {
		log.Println("error calling taskRepository.getById() method: ", err)
		err := tx.Rollback()
		if err != nil {
			log.Println("error while rolling back transaction in taskRepository.getById() method : ", err)
		}
		return nil, err
	}

	for _, headerFromInput := range taskEntity.ResponseHeaders {
		for _, headerFromDb := range *responseHeaders {
			if strings.EqualFold(headerFromInput.Name, headerFromDb.Name) {

			}
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println("error while commiting transaction in taskRepository.getById() method : ", err)
		return nil, err
	}

	var taskDto = c.TaskMapper.MapToDto(*resultTask)
	return &taskDto, nil
}

func (c TaskService) DeleteById(id string) error {
	tx, err := database.DB.Begin()
	if err != nil {
		log.Println("error while opening transaction taskRepository.getById() method: ", err)
		return err
	}
	err = c.TaskRepository.DeleteById(tx, id)

	if err != nil {
		log.Println("error calling taskRepository.getById() method: ", err)
		err := tx.Rollback()
		if err != nil {
			log.Println("error while rolling back transaction in taskRepository.getById() method : ", err)
		}
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

func (tm TaskMapper) MapToDto(task domain.Task) dto.TaskDTO {
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

	for _, header := range task.RequestHeaders {
		taskDto.RequestHeaders[header.Name] = header.Value
	}

	for _, header := range task.ResponseHeaders {
		taskDto.ResponseHeaders[header.Name] = header.Value
	}

	return taskDto
}

func (tm TaskMapper) MapToEntity(taskDTO dto.TaskDTO) domain.Task {
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

	return task
}

// func ErrorHandler(fn Executable) error {
// 	err := fn()
// 	if err != nil {
// 		defer func() {
// 			if p := recover(); p != nil {
// 				// a panic occurred, rollback and repanic
// 				tx.Rollback()
// 				panic(p)
// 			} else if err != nil {
// 				// something went wrong, rollback
// 				tx.Rollback()
// 			} else {
// 				// all good, commit
// 				err = tx.Commit()
// 			}
// 		}()

// 	}
// 	return err
// }

// type Executable func() error
