package repository

import (
	"database/sql"
	"log"

	"example.com/test_axxonsoft/v2/domain"
	"github.com/google/uuid"
	sqlbuilder "github.com/huandu/go-sqlbuilder"
)

type TaskRepository struct {
}

// Gets existing task entity from DB by id
// if not found returns nil
func (t TaskRepository) GetById(tx *sql.Tx, id string) (*domain.Task, error) {
	sb := sqlbuilder.NewSelectBuilder()

	sb.Select("id", "method", "url", "http_status_code", "task_status", "response_length",
		"request_body", "response_body").From("task").Where(sb.Equal("id", id))
	query, args := sb.Build()
	rows, err := tx.Query(query, args)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return nil, err
	}

	task := new(domain.Task)
	if rows.Next() {

		err = rows.Scan(&task.Id, &task.Method, &task.Url,
			&task.HttpStatusCode, &task.TaskStatus, &task.ResponseLength)
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

// Creates new task entity
// returns new create task entity , with unique id
func (t TaskRepository) Create(tx *sql.Tx, task *domain.Task) (*domain.Task, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		log.Println("an error occurred while generating uuid : ", err.Error())
		return nil, err
	}

	task.Id = uid.String()

	sb := sqlbuilder.NewInsertBuilder()
	sb.InsertInto("task").Cols("id", "method", "url", "http_status_code", "task_status", "response_length",
		"request_body", "response_body").Values(task.Id, task.Method, task.Url,
		task.HttpStatusCode, task.TaskStatus, task.ResponseLength)
	query, args := sb.Build()
	_, err = tx.Exec(query, args)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return nil, err
	}
	return task, nil
}

// Updates task entity
// returns Updated task entity
func (t TaskRepository) Update(tx *sql.Tx, task *domain.Task) (*domain.Task, error) {
	sb := sqlbuilder.NewUpdateBuilder()
	sb.Update("task").
		Set(sb.Equal("id", task.Id)).
		Set(sb.Equal("method", task.Method)).
		Set(sb.Equal("url", task.Url)).
		Set(sb.Equal("http_status_code", task.HttpStatusCode)).
		Set(sb.Equal("task_status", task.TaskStatus)).
		Set(sb.Equal("response_length", task.ResponseLength)).
		Set(sb.Equal("request_body", task.RequestBody)).
		Set(sb.Equal("response_body", task.ResponseBody)).
		Where(sb.Equal("id", task.Id))
	query, args := sb.Build()
	_, err := tx.Exec(query, args)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return nil, err
	}
	return task, nil
}

// Gets all tasks from DB with pagination
// page represents number of page in DB, starts from 0
// size represents size of the page fetched from DB
func (t TaskRepository) FindAll(tx *sql.Tx, page, size int) (*[]domain.Task, error) {
	sb := sqlbuilder.NewSelectBuilder()

	sb.Select("id", "method", "url", "http_status_code", "task_status", "response_length",
		"request_body", "response_body").From("task").Offset(page * size).Limit(size)
	query, args := sb.Build()
	rows, err := tx.Query(query, args)
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

// deletes task entity  in DB by task
func (t TaskRepository) DeleteById(tx *sql.Tx, id string) error {
	sb := sqlbuilder.NewDeleteBuilder()
	sb.DeleteFrom("task").Where(sb.Equal("id", id))
	query, args := sb.Build()
	_, err := tx.Exec(query, args)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return err
	}
	return nil
}
