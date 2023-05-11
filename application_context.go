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
	appContext.RabbitContext = rabbitContext
	appContext.TaskService = &taskService
	var headerRepository = repository.HeaderRepository{}
	var taskRepository = repository.TaskRepository{}
	appContext.TaskService.HeaderRepository = &headerRepository
	appContext.TaskService.TaskRepository = &taskRepository
	var taskController = controller.TaskController{}
	taskController.RabbitContext = appContext.RabbitContext
	taskController.TaskService = &taskService
	appContext.TaskController = &taskController
	appContext.TaskService.RabbitContext = appContext.RabbitContext
	return &appContext
}

func (a *ApplicationContext) Close() {
	a.RabbitContext.Consumer.Close()
	a.RabbitContext.Publisher.Close()
}
