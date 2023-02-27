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
}

func (c TaskService) GetById(id string) (*dto.TaskDTO, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		log.Println("error while opening transaction taskRepository.getById() method: ", err)
		return nil, err
	}
	task, err := c.TaskRepository.GetById(tx, id)

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

	var taskDto = dto.TaskDTO{
		Id:             task.Id,
		Method:         task.Method,
		Url:            task.Url,
		HttpStatusCode: task.HttpStatusCode,
		ResponseLength: task.ResponseLength,
		TaskStatus:     task.TaskStatus,
		RequestBody:    task.RequestBody,
		ResponseBody:   task.ResponseBody,
	}
	return &taskDto, nil
}

func (c TaskService) Create(taskDTO dto.TaskDTO) (*dto.TaskDTO, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		log.Println("error while opening transaction taskRepository.getById() method: ", err)
		return nil, err
	}
	resultTask, err := c.TaskRepository.Create(tx, id)

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

	var taskDto = dto.TaskDTO{
		Id:             resultTask.Id,
		Method:         resultTask.Method,
		Url:            resultTask.Url,
		HttpStatusCode: resultTask.HttpStatusCode,
		ResponseLength: resultTask.ResponseLength,
		TaskStatus:     resultTask.TaskStatus,
		RequestBody:    resultTask.RequestBody,
		ResponseBody:   resultTask.ResponseBody,
	}
	return &taskDto, nil
}

func (c TaskService) Update(taskDTO dto.TaskDTO) (*dto.TaskDTO, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		log.Println("error while opening transaction taskRepository.getById() method: ", err)
		return nil, err
	}
	task := domain.Task{}
	taskResult, err := c.TaskRepository.Update(tx, &task)

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

	var taskDto = dto.TaskDTO{
		Id:             taskResult.Id,
		Method:         taskResult.Method,
		Url:            taskResult.Url,
		HttpStatusCode: taskResult.HttpStatusCode,
		ResponseLength: taskResult.ResponseLength,
		TaskStatus:     taskResult.TaskStatus,
		RequestBody:    taskResult.RequestBody,
		ResponseBody:   taskResult.ResponseBody,
	}
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
