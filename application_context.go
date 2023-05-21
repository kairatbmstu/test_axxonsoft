package main

import (
	"example.com/test_axxonsoft/v2/controller"
	"example.com/test_axxonsoft/v2/repository"
	"example.com/test_axxonsoft/v2/service"
)

type PostrgresConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type RabbitMqConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type ApplicationContext struct {
	TaskController *controller.TaskController
	TaskService    *service.TaskService
	RabbitContext  *service.RabbitContext
}

func NewApplicationContext() *ApplicationContext {
	var appContext = ApplicationContext{}
	var taskService = service.TaskService{}

	var headerRepository = repository.HeaderRepository{}
	var taskRepository = repository.TaskRepository{}
	var taskController = controller.TaskController{}

	appContext.TaskService = &taskService
	appContext.TaskService.HeaderRepository = &headerRepository
	appContext.TaskService.TaskRepository = &taskRepository

	taskController.TaskService = &taskService
	appContext.TaskController = &taskController

	var rabbitContext = service.NewRabbitContext()
	appContext.RabbitContext = rabbitContext
	appContext.TaskService.RabbitContext = rabbitContext
	appContext.RabbitContext.TaskService = &taskService
	rabbitContext.Init()
	return &appContext
}

func (a *ApplicationContext) Close() {
	a.RabbitContext.Consumer.Close()
	a.RabbitContext.Publisher.Close()
}
