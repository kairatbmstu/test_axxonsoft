package app

import (
	"example.com/test_axxonsoft/v2/repository"
	"example.com/test_axxonsoft/v2/service"
)

type ApplicationContext struct {
	taskService      *service.TaskService
	taskRepository   *repository.TaskRepository
	headerRepository repository.HeaderRepository
}

func InitApplicationContext() *ApplicationContext {
	var appContext = new(ApplicationContext)
	return appContext
}
