package main

import (
	"example.com/test_axxonsoft/v2/controller"
	"example.com/test_axxonsoft/v2/repository"
	"example.com/test_axxonsoft/v2/service"
)

type ApplicationContext struct {
	TaskController *controller.TaskController
	TaskService    *service.TaskService
	RabbitContext  *service.RabbitContext
}

func BuildApplicationContext() *ApplicationContext {
	var appContext = ApplicationContext{}
	var taskService = service.TaskService{}
	var rabbitContext = service.InitRabbitContext()
	var headerRepository = repository.HeaderRepository{}
	var taskRepository = repository.TaskRepository{}
	var taskController = controller.TaskController{}

	appContext.RabbitContext = rabbitContext
	appContext.TaskService = &taskService
	appContext.TaskService.HeaderRepository = &headerRepository
	appContext.TaskService.TaskRepository = &taskRepository
	taskController.RabbitContext = rabbitContext
	taskController.TaskService = &taskService
	appContext.TaskController = &taskController
	appContext.TaskService.RabbitContext = rabbitContext
	appContext.RabbitContext.TaskService = &taskService
	return &appContext
}

func (a *ApplicationContext) Close() {
	a.RabbitContext.Consumer.Close()
	a.RabbitContext.Publisher.Close()
}
