package repository

import (
	"database/sql"
	"log"

	"example.com/test_axxonsoft/v2/domain"
	"github.com/google/uuid"
)

const (
	GET_TASK    = `select * from task where id = $1`
	INSERT_TASK = `insert into task() values()`
	UPDATE_TASK = `update task set `
	DELETE_TASK = `delete from task where id = $1`
	GET_TASKS   = `select * from task limit $1 offset $2`
)

type TaskRepository struct {
}

func (t TaskRepository) GetById(tx *sql.Tx, id string) (*domain.Task, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		log.Println("an error occurred while generating uuid : ", err.Error())
		return nil, err
	}
	header.Id = uid.String()
	_, err = tx.Exec(insertHeaderStmt, header.Id, header.Name, header.RequestTaskId, header.ResponsetTaskId, header.Value)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return nil, err
	}
	return header, nil
}

func (t TaskRepository) Save(tx *sql.Tx, task *domain.Task) (*domain.Task, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		log.Println("an error occurred while generating uuid : ", err.Error())
		return nil, err
	}
	header.Id = uid.String()
	_, err = tx.Exec(INSERT_TASK, header.Id, header.Name, header.RequestTaskId, header.ResponsetTaskId, header.Value)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return nil, err
	}
	return header, nil
}

func (t TaskRepository) Update(tx *sql.Tx, task *domain.Task) (*domain.Task, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		log.Println("an error occurred while generating uuid : ", err.Error())
		return nil, err
	}
	header.Id = uid.String()
	_, err = tx.Exec(UPDATE_TASK, header.Id, header.Name, header.RequestTaskId, header.ResponsetTaskId, header.Value)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return nil, err
	}
	return header, nil
}

func (t TaskRepository) FindAll(tx *sql.Tx, page, size int) (*[]domain.Task, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		log.Println("an error occurred while generating uuid : ", err.Error())
		return nil, err
	}
	header.Id = uid.String()
	_, err = tx.Exec(GET_TASKS, header.Id, header.Name, header.RequestTaskId, header.ResponsetTaskId, header.Value)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return nil, err
	}
	return header, nil
}

func (t TaskRepository) DeleteById(tx *sql.Tx, id string) error {
	_, err := tx.Exec(DELETE_TASK, id)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return err
	}
	return nil
}
