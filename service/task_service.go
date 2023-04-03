package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"example.com/test_axxonsoft/v2/database"
	"example.com/test_axxonsoft/v2/domain"
	"example.com/test_axxonsoft/v2/dto"
	"example.com/test_axxonsoft/v2/rabbit"
	"example.com/test_axxonsoft/v2/repository"
	"github.com/gofrs/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
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

// Create New Task entity
func (c TaskService) CreateNewTask(taskDTO dto.TaskDTO) (*dto.TaskDTO, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		log.Println("error while opening transaction taskRepository.getById() method: ", err)
		return nil, err
	}

	taskDTO.TaskStatus = domain.TaskStatus_NEW
	task := c.TaskMapper.MapToEntity(taskDTO)
	err = c.TaskRepository.Create(tx, &task)

	if err != nil {
		log.Println("error calling taskRepository.getById() method: ", err)
		err := tx.Rollback()
		if err != nil {
			log.Println("error while rolling back transaction in taskRepository.getById() method : ", err)
		}
		return nil, err
	}

	for _, header := range task.RequestHeaders {
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
		log.Println("error while commiting transaction in taskRepository.CreateNewTask() method : ", err)
		return nil, err
	}

	var taskDto = c.TaskMapper.MapToDto(task)

	err = c.SendToQueue(taskDto)

	if err != nil {
		log.Println("error while sending  taskDto in tasks-queue : ", err)
		return nil, err
	}

	return &taskDto, nil
}

func (c TaskService) SendToQueue(taskDTO dto.TaskDTO) error {
	conn, err := rabbit.GetConnection()
	if err != nil {
		return err
	}
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	defer conn.Close()

	queue, err := ch.QueueDeclare(
		"tasks-queue", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Println("  : ", err)
		return nil
	}

	err = ch.ExchangeDeclare(
		"tasks-exchange", // name
		"direct",         // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)

	if err != nil {
		log.Println(" : ", err)
		return nil
	}

	err = ch.QueueBind(
		queue.Name,       // queue name
		"",               // routing key
		"tasks-exchange", // exchange
		false,
		nil)

	if err != nil {
		log.Println(" : ", err)
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(taskDTO)
	if err != nil {
		return nil
	}
	err = ch.PublishWithContext(ctx,
		"tasks-exchange", // exchange
		queue.Name,       // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		log.Println("error while commiting transaction in taskRepository.getById() method : ", err)
		return nil
	}
	log.Printf(" [x] Sent %s\n", body)

	// save and change task status to in_process
	// send taskDto to Queue
	return nil
}

func (c TaskService) ReceiveFromQueue(message string) error {
	//receive taskDto
	//make http request
	//handle response from http request
	//if http ~ 4xx or 5xx
	//save response with task status error
	//if http status ~ 2xx
	//save response with task status done
	return nil
}

// Saves Response received from http request
func (c TaskService) SaveResponse(taskDTO dto.TaskDTO) (*dto.TaskDTO, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		log.Println("error while opening transaction taskRepository.getById() method: ", err)
		return nil, err
	}

	taskEntity := c.TaskMapper.MapToEntity(taskDTO)
	err = c.TaskRepository.Update(tx, &taskEntity)

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
	return &taskDto, nil
}

func (c TaskService) ChangeTaskStatus(taskDTO dto.TaskDTO) error {

	tx, err := database.DB.Begin()
	if err != nil {
		log.Println("error while opening transaction taskRepository.getById() method: ", err)
		return err
	}

	taskEntity := c.TaskMapper.MapToEntity(taskDTO)
	err = c.TaskRepository.ChangeTaskStatus(tx, &taskEntity)

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

	taskDto.RequestHeaders = make(map[string]string)
	for _, header := range task.RequestHeaders {
		taskDto.RequestHeaders[header.Name] = header.Value
	}

	taskDto.ResponseHeaders = make(map[string]string)
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
