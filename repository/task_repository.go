package repository

import (
	"database/sql"
	"fmt"
	"log"

	"example.com/test_axxonsoft/v2/database"
	"example.com/test_axxonsoft/v2/domain"
	"github.com/google/uuid"
	sqlbuilder "github.com/huandu/go-sqlbuilder"
)

const (
	TASK_UPDATE = `UPDATE task SET method=$1, 
					 url=$2, 
					 http_status_code=$3,
					 task_status=$4, 
					 response_length=$5, 
					 request_body=$6, 
					 response_body=$7 WHERE id = $8`
)

/*
The TaskRepository type represents a repository for managing task entities in the database.
*/
type TaskRepository struct {
}

/**
The GetById method retrieves an existing task entity from the database by its ID. If the task is not found, it returns nil.

Parameters
tx: A transaction object for executing the database query.
id: The ID of the task entity to retrieve.
Returns
*domain.Task: The retrieved task entity.
error: An error object if an error occurred during the retrieval process.
*/
func (t TaskRepository) GetById(tx *sql.Tx, id uuid.UUID) (*domain.Task, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "method", "url", "http_status_code", "task_status", "response_length",
		"request_body", "response_body").From("task").Where(sb.Equal("id", id))
	query, args := sb.Build()
	rows, err := tx.Query(query, args...)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return nil, err
	}

	defer rows.Close()

	task := new(domain.Task)
	if rows.Next() {

		err = rows.Scan(&task.Id, &task.Method, &task.Url,
			&task.HttpStatusCode, &task.TaskStatus, &task.ResponseLength, &task.RequestBody, &task.ResponseBody)
		if err != nil {
			log.Println("err : ", err)
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
	}

	return task, nil
}

/**
The GetTaskStatusById method retrieves the task status of an existing task entity from the database by its ID. If the task is not found, it returns nil.

Parameters
tx: A transaction object for executing the database query.
id: The ID of the task entity to retrieve the status.
Returns
*domain.Task: The retrieved task entity with only the task status field populated.
error: An error object if an error occurred during the retrieval process.
*/
func (t TaskRepository) GetTaskStatusById(tx *sql.Tx, id uuid.UUID) (*domain.Task, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "method", "url", "http_status_code", "task_status", "response_length",
		"request_body").From("task").Where(sb.Equal("id", id))
	query, args := sb.Build()
	rows, err := tx.Query(query, args...)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return nil, err
	}

	defer rows.Close()

	task := new(domain.Task)
	if rows.Next() {

		err = rows.Scan(&task.Id, &task.Method, &task.Url,
			&task.HttpStatusCode, &task.TaskStatus, &task.ResponseLength, &task.RequestBody)
		if err != nil {
			log.Println("err : ", err)
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
	}

	return task, nil
}

/*
The Create method creates a new task entity in the database.

Parameters
task: The task entity to create.
Returns
error: An error object if an error occurred during the creation process.
*/
func (t TaskRepository) Create(task *domain.Task) error {
	uid := uuid.New()

	log.Println("create entity with uid : ", uid)
	task.Id = uid

	for i := range task.RequestHeaders {
		task.RequestHeaders[i].RequestTaskId = &task.Id
	}

	sb := sqlbuilder.PostgreSQL.NewInsertBuilder()
	sb.InsertInto("task").Cols("id", "method", "url", "http_status_code", "task_status", "response_length",
		"request_body", "response_body").Values(task.Id.String(), task.Method, task.Url,
		task.HttpStatusCode, task.TaskStatus, task.ResponseLength, task.RequestBody, task.ResponseBody)
	query, args := sb.Build()
	_, err := database.DB.Exec(query, args...)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return err
	}

	return nil
}

/*
The Update method updates an existing task entity in the database.

Parameters
task: The task entity to update.
Returns
error: An error object if an error occurred during the update process.
*/
func (t TaskRepository) Update(task *domain.Task) error {
	log.Println("Update task : ", task.Id.String())
	for i := range task.ResponseHeaders {
		task.ResponseHeaders[i].ResponseTaskId = &task.Id
	}

	//sb := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	// sb.Update("task").
	// 	Set(sb.Assign("method", task.Method),
	// 		sb.Assign("url", task.Url),
	// 		sb.Assign("http_status_code", task.HttpStatusCode),
	// 		sb.Assign("task_status", task.TaskStatus),
	// 		sb.Assign("response_length", task.ResponseLength),
	// 		sb.Assign("request_body", task.RequestBody),
	// 		sb.Assign("response_body", task.ResponseBody),
	// 	).
	// 	Where(sb.Equal("id", task.Id))
	// query, args := sb.Build()
	_, err := database.DB.Exec(TASK_UPDATE, task.Method, task.Url, task.HttpStatusCode, task.TaskStatus,
		task.ResponseLength, task.RequestBody, task.ResponseBody, task.Id.String())
	if err != nil {
		log.Println("an error occurred while executing update statement : ", err.Error())
		return err
	}
	return nil
}

/*
The ChangeTaskStatus method changes the status of a task entity in the database.

Parameters
taskId: The ID of the task entity to update the status.
taskStatus: The new task status value.
Returns
error: An error object if an error occurred during the status change process.
*/
func (t TaskRepository) ChangeTaskStatus(taskId uuid.UUID, taskStatus domain.TaskStatus) error {
	log.Println("changeTaskStatus :  ", taskId.String())
	sb := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	sb.Update("task")
	sb.Set(sb.Assign("task_status", taskStatus))
	sb.Where(sb.Equal("id", taskId.String()))
	query, args := sb.Build()
	fmt.Println(query)
	_, err := database.DB.Exec(query, args...)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return err
	}
	return nil
}

/*
The FindAll method retrieves all tasks from the database with pagination.

Parameters
tx: A transaction object for executing the database query.
page: The number of the page to retrieve (starting from 0).
size: The size of the page (number of tasks per page).
Returns
*[]domain.Task: A pointer to a slice containing the retrieved task entities.
error: An error object if an error occurred during the retrieval process.
*/
func (t TaskRepository) FindAll(tx *sql.Tx, page, size int) (*[]domain.Task, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "method", "url", "http_status_code", "task_status", "response_length",
		"request_body", "response_body").From("task").Offset(page * size).Limit(size)
	query, args := sb.Build()
	rows, err := tx.Query(query, args...)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return nil, err
	}

	result := make([]domain.Task, 0)

	for rows.Next() {
		task := new(domain.Task)
		err = rows.Scan(&task.Id, &task.Method, &task.Url,
			&task.HttpStatusCode, &task.TaskStatus, &task.ResponseLength)
		if err != nil {
			log.Println("err : ", err)
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		result = append(result, *task)
	}

	return &result, nil
}

/*
The DeleteById method deletes a task entity from the database by its ID.

Parameters
tx: A transaction object for executing the database query.
id: The ID of the task entity to delete.
Returns
error: An error object if an error occurred during the deletion process.
*/
func (t TaskRepository) DeleteById(tx *sql.Tx, id string) error {
	sb := sqlbuilder.PostgreSQL.NewDeleteBuilder()
	sb.DeleteFrom("task").Where(sb.Equal("id", id))
	query, args := sb.Build()
	_, err := tx.Exec(query, args...)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return err
	}
	return nil
}
