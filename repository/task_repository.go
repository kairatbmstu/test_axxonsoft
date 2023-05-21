package repository

import (
	"database/sql"
	"log"

	"example.com/test_axxonsoft/v2/domain"
	"github.com/gofrs/uuid"
	sqlbuilder "github.com/huandu/go-sqlbuilder"
)

const (
	TASK_UPDATE = `UPDATE task SET method = $1, url = $2, http_status_code = $3, task_status = $4, response_length = $5, request_body = $6, response_body = $7 WHERE id = $8`
)

type TaskRepository struct {
}

// Gets existing task entity from DB by id
// if not found returns nil
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

// Creates new task entity
// returns new create task entity , with unique id
func (t TaskRepository) Create(tx *sql.Tx, task *domain.Task) error {
	uid, err := uuid.NewV7()
	if err != nil {
		log.Println("an error occurred while generating uuid : ", err.Error())
		return err
	}

	task.Id = uid

	for i := range task.RequestHeaders {
		task.RequestHeaders[i].RequestTaskId = &task.Id
	}

	sb := sqlbuilder.PostgreSQL.NewInsertBuilder()
	sb.InsertInto("task").Cols("id", "method", "url", "http_status_code", "task_status", "response_length",
		"request_body", "response_body").Values(task.Id, task.Method, task.Url,
		task.HttpStatusCode, task.TaskStatus, task.ResponseLength, task.RequestBody, task.ResponseBody)
	query, args := sb.Build()
	_, err = tx.Exec(query, args...)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return err
	}

	return nil
}

// Updates task entity
// returns Updated task entity
func (t TaskRepository) Update(tx *sql.Tx, task *domain.Task) error {

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
	_, err := tx.Exec(TASK_UPDATE, task.Method, task.Url, task.HttpStatusCode, task.TaskStatus,
		task.ResponseLength, task.RequestBody, "task.ResponseBody", task.Id)
	if err != nil {
		log.Println("an error occurred while executing update statement : ", err.Error())
		return err
	}
	return nil
}

// Change task status
func (t TaskRepository) ChangeTaskStatus(tx *sql.Tx, taskId uuid.UUID, taskStatus domain.TaskStatus) error {
	sb := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	sb.Update("task").
		Set(sb.Assign("task_status", taskStatus)).
		Where(sb.Equal("id", taskId))
	query, args := sb.Build()
	_, err := tx.Exec(query, args...)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return err
	}
	return nil
}

// Gets all tasks from DB with pagination
// page represents number of page in DB, starts from 0
// size represents size of the page fetched from DB
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

// deletes task entity  in DB by task
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
