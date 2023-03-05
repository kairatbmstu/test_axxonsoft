package repository

import (
	"database/sql"
	"log"

	"example.com/test_axxonsoft/v2/domain"
	"github.com/google/uuid"
)

const (
	INSERT_HEADER        = `insert into headers(id,request_headers_task_id,response_headers_task_id,header_name,header_value) values($1, $2, $3, $4)`
	UPDATE_HEADER        = `update headers set where id = $1`
	DELETE_HEADER        = `delete from headers where id = $1`
	GET_HEADER           = `select * from headers where id = $1`
	GET_REQUEST_HEADERS  = `select * from headers where request_headers_task_id = $1`
	GET_RESPONSE_HEADERS = `select * from headers where response_headers_task_id = $1`
)

type HeaderRepository struct {
}

func (h HeaderRepository) Save(tx *sql.Tx, header *domain.Header) (*domain.Header, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		log.Println("an error occurred while generating uuid : ", err.Error())
		return nil, err
	}
	header.Id = uid.String()
	_, err = tx.Exec(INSERT_HEADER, header.Id, header.Name, header.RequestTaskId, header.ResponsetTaskId, header.Value)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return nil, err
	}
	return header, nil
}

func (h HeaderRepository) Update(tx *sql.Tx, header *domain.Header) (*domain.Header, error) {
	_, err := tx.Exec(UPDATE_HEADER, header.Id, header.Name, header.RequestTaskId, header.ResponsetTaskId, header.Value)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return nil, err
	}
	return header, nil
}

func (h HeaderRepository) GetRequestHeaders(tx *sql.Tx, taskId string) (*[]domain.Header, error) {
	rows, err := tx.Query(GET_REQUEST_HEADERS, taskId)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return nil, err
	}

	var result = make([]domain.Header, 0)

	for rows.Next() {
		rows.Scan()
	}

	return &result, nil
}

func (h HeaderRepository) GetResponseHeaders(tx *sql.Tx, taskId string) (*[]domain.Header, error) {
	rows, err := tx.Query(GET_REQUEST_HEADERS, taskId)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return nil, err
	}

	var result = make([]domain.Header, 0)

	for rows.Next() {
		rows.Scan()
	}

	return &result, nil
}
