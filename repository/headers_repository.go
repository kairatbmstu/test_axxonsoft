package repository

import (
	"database/sql"
	"log"

	"example.com/test_axxonsoft/v2/domain"
	"github.com/google/uuid"
)

const (
	insertHeaderStmt = `insert into headers(id,request_headers_task_id,response_headers_task_id,header_name,header_value) values($1, $2, $3, $4)`
)

func SaveHeader(tx *sql.Tx, header *domain.Header) (*domain.Header, error) {
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

func UpdateHeader(tx *sql.Tx, header *domain.Header) (*domain.Header, error) {
	return nil, nil
}

func GetRequestHeaders(tx *sql.Tx, taskId string) (*[]domain.Header, error) {
	return nil, nil
}

func GetResponseHeaders(tx *sql.Tx, taskId string) (*[]domain.Header, error) {
	return nil, nil
}
