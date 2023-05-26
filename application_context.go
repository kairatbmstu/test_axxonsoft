package main

import (
	"example.com/test_axxonsoft/v2/config"
	"example.com/test_axxonsoft/v2/controller"
	"example.com/test_axxonsoft/v2/repository"
	"example.com/test_axxonsoft/v2/service"
)

type ApplicationContext struct {
	TaskController *controller.TaskController
	TaskService    *service.TaskService
	RabbitContext  *service.RabbitContext
}

func NewApplicationContext(rabbitmqConfig config.RabbitMqConfig) *ApplicationContext {
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

	var rabbitContext = service.NewRabbitContext(rabbitmqConfig)

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
