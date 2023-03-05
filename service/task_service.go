package service

import (
	"log"

	"example.com/test_axxonsoft/v2/database"
	"example.com/test_axxonsoft/v2/domain"
	"example.com/test_axxonsoft/v2/dto"
	"example.com/test_axxonsoft/v2/repository"
)

type TaskService struct {
	TaskRepository   repository.TaskRepository
	HeaderRepository repository.HeaderRepository
	TaskMapper       TaskMapper
}

func (t TaskService) GetById(id string) (*dto.TaskDTO, error) {
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
		log.Println("error calling taskRepository.getById() method: ", err)
		err := tx.Rollback()
		if err != nil {
			log.Println("error while rolling back transaction in taskRepository.getById() method : ", err)
		}
		return nil, err
	}

	task.RequestHeaders = *requestHeaders

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
	task := domain.Task{}
	resultTask, err := c.TaskRepository.Update(tx, &task)

	if err != nil {
		log.Println("error calling taskRepository.getById() method: ", err)
		err := tx.Rollback()
		if err != nil {
			log.Println("error while rolling back transaction in taskRepository.getById() method : ", err)
		}
		return nil, err
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
		var headerDto = dto.HeaderDTO{
			Name:  header.Name,
			Value: header.Value,
		}
		taskDto.RequestHeaders = append(taskDto.RequestHeaders, headerDto)
	}

	for _, header := range task.ResponseHeaders {
		var headerDto = dto.HeaderDTO{
			Name:  header.Name,
			Value: header.Value,
		}
		taskDto.ResponseHeaders = append(taskDto.ResponseHeaders, headerDto)
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

	for _, headerDto := range taskDTO.ResponseHeaders {
		var header = domain.Header{
			Name:  headerDto.Name,
			Value: headerDto.Value,
		}
		task.ResponseHeaders = append(task.ResponseHeaders, header)
	}

	return task
}
